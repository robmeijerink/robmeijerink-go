export const initDataHighway = () => {
    const canvas = document.getElementById('flow-canvas');
    if (!canvas) return;
    const ctx = canvas.getContext('2d');

    let width, height, dpr, tracers = [], pulses = [];
    let highwayMinY = 0;
    let highwayMaxY = 0;

    const accentColor = '#B84A2B';
    const gridSize = 24;

    const resize = () => {
        dpr = window.devicePixelRatio || 1;
        width = window.innerWidth;
        height = window.innerHeight;

        canvas.width = width * dpr;
        canvas.height = height * dpr;
        canvas.style.width = width + 'px';
        canvas.style.height = height + 'px';

        ctx.scale(dpr, dpr);
        updateHighwayBounds();
    };

    const updateHighwayBounds = () => {
        const heroSection = document.getElementById('hero-section');

        if (heroSection) {
            const heroRect = heroSection.getBoundingClientRect();
            const heroBottom = heroRect.bottom + window.scrollY;

            highwayMinY = 0;

            highwayMaxY = Math.floor(heroBottom / gridSize) * gridSize;
        }
    };

    class Pulse {
        constructor(x, y) {
            this.x = x;
            this.y = y;
            this.r = 1;
            this.opacity = 0.5;
        }
        update() {
            this.r += 1.5;
            this.opacity -= 0.02;
        }
        draw() {
            const renderY = this.y - window.scrollY;
            if (renderY < -50 || renderY > height + 50) return;

            ctx.beginPath();
            ctx.arc(this.x, renderY, this.r, 0, Math.PI * 2);
            ctx.strokeStyle = accentColor;
            ctx.lineWidth = 1;
            ctx.globalAlpha = Math.max(0, this.opacity);
            ctx.stroke();
        }
    }

    /**
     * Tracer Class - Handles the movement and drawing of individual data lines.
     * Strictly enforces grid alignment and 90-degree turns.
     */
    class Tracer {
        constructor() {
            this.reset();
        }

        reset() {
            this.speed = 1.2;
            // Ensure bounds are defined, fallback to window innerHeight
            const min = typeof highwayMinY !== 'undefined' ? highwayMinY : 0;
            const max = typeof highwayMaxY !== 'undefined' ? highwayMaxY : window.innerHeight;

            // Snap initial Y to grid
            this.y = Math.floor((Math.random() * (max - min) + min) / gridSize) * gridSize;

            // Start from left or right edge
            this.x = (Math.random() > 0.5) ? -gridSize : width + gridSize;

            // Initial velocity (horizontal only)
            this.vx = (this.x < 0) ? this.speed : -this.speed;
            this.vy = 0;

            this.history = [];
            this.maxLength = 22;
            this.distMoved = 0;
            this.turnTarget = this.getRandomTurnDistance();
            this.pulseCooldown = 0;
        }

        getRandomTurnDistance() {
            // Random turn distance between 2 and 7 grid blocks
            return (Math.floor(Math.random() * 6) + 2) * gridSize;
        }

        update() {
            // Apply velocity
            this.x += this.vx;
            this.y += this.vy;
            this.distMoved += this.speed;

            const min = typeof highwayMinY !== 'undefined' ? highwayMinY : 0;
            const max = typeof highwayMaxY !== 'undefined' ? highwayMaxY : window.innerHeight;

            // 1. Handle Turn Logic (90 degrees only)
            if (this.distMoved >= this.turnTarget) {
                // Snap to grid
                this.x = Math.round(this.x / gridSize) * gridSize;
                this.y = Math.round(this.y / gridSize) * gridSize;

                // Switch axis
                if (this.vx !== 0) {
                    this.vx = 0;
                    if (this.y <= min) {
                        this.vy = this.speed;
                    } else if (this.y >= max) {
                        this.vy = -this.speed;
                    } else {
                        this.vy = (Math.random() > 0.5 ? this.speed : -this.speed);
                    }
                } else {
                    this.vy = 0;
                    this.vx = (Math.random() > 0.5 ? this.speed : -this.speed);
                }

                this.distMoved = 0;
                this.turnTarget = this.getRandomTurnDistance();
            }

            // 2. Bound checking (Bounce back to stay within area)
            // Fix: Only trigger if explicitly moving out of bounds
            if (this.y <= min && this.vy < 0) {
                this.y = min;
                this.vy = 0;
                this.vx = (Math.random() > 0.5 ? this.speed : -this.speed);
                this.distMoved = 0;
            } else if (this.y >= max && this.vy > 0) {
                this.y = max;
                this.vy = 0;
                this.vx = (Math.random() > 0.5 ? this.speed : -this.speed);
                this.distMoved = 0;
            }

            // Update history for tail
            this.history.push({x: this.x, y: this.y});
            if (this.history.length > this.maxLength) this.history.shift();

            // 3. Reset if completely off-screen
            if (this.x < -200 || this.x > width + 200) this.reset();
        }

        draw() {
            if (this.history.length < 2) return;

            const renderY = this.y - window.scrollY;
            // Only draw if within vertical view
            if (renderY < -100 || renderY > height + 100) return;

            // Draw trail
            ctx.beginPath();
            ctx.strokeStyle = accentColor;
            ctx.lineWidth = 1.5;
            ctx.globalAlpha = 0.4;
            ctx.moveTo(this.history[0].x, this.history[0].y - window.scrollY);

            for(let i = 1; i < this.history.length; i++) {
                ctx.lineTo(this.history[i].x, this.history[i].y - window.scrollY);
            }
            ctx.stroke();

            // Draw head
            ctx.beginPath();
            ctx.arc(this.x, renderY, 2.5, 0, Math.PI * 2);
            ctx.fillStyle = accentColor;
            ctx.globalAlpha = 0.8;
            ctx.fill();
        }
    }

    /**
     * Main animation loop
     */
    const animate = () => {
        ctx.clearRect(0, 0, width, height);

        // 1. Draw background grid
        ctx.beginPath();
        ctx.strokeStyle = accentColor;
        ctx.globalAlpha = 0.04;
        ctx.lineWidth = 1;

        // Vertical lines
        for(let i = 0; i < width; i += gridSize) {
            ctx.moveTo(i, 0);
            ctx.lineTo(i, height);
        }

        // Horizontal lines (scrolled offset)
        const scrollOffset = window.scrollY % gridSize;
        for(let j = -scrollOffset; j < height + gridSize; j += gridSize) {
            ctx.moveTo(0, j);
            ctx.lineTo(width, j);
        }
        ctx.stroke();

        // 2. Detect Data Handshakes (collision logic)
        for (let i = 0; i < tracers.length; i++) {
            for (let j = i + 1; j < tracers.length; j++) {
                const t1 = tracers[i];
                const t2 = tracers[j];

                // Check proximity on grid
                if (Math.abs(t1.x - t2.x) < 10 && Math.abs(t1.y - t2.y) < 10) {
                    // If cooldown is expired, trigger handshake pulse
                    if (!t1.pulseCooldown || t1.pulseCooldown <= 0) {
                        pulses.push(new Pulse(t1.x, t1.y));

                        t1.pulseCooldown = 40;
                        t2.pulseCooldown = 40;
                    }
                }
            }
            // Decrease pulse cooldowns
            if (tracers[i].pulseCooldown > 0) tracers[i].pulseCooldown--;
        }

        // 3. Update and draw pulses
        pulses.forEach((p, i) => {
            p.update();
            p.draw();
            if (p.opacity <= 0) pulses.splice(i, 1);
        });

        // 4. Update and draw tracers
        tracers.forEach(t => {
            t.update();
            t.draw();
        });

        requestAnimationFrame(animate);
    };

    const initEngine = () => {
        resize();
        // Reduced from 8 to 6 tracers to prioritize visual hierarchy and breathing room
        for (let i = 0; i < 6; i++) {
            tracers.push(new Tracer());
        }
        animate();
    };

    window.addEventListener('resize', resize);
    setTimeout(updateHighwayBounds, 500);

    initEngine();
};

export const initBlueprintSteps = () => {
   // Target the inner container instead of the whole section to delay the start point
    const container = document.getElementById('blueprint-steps-container');
    const progressBar = document.getElementById('blueprint-progress-bar');
    const step1 = document.getElementById('blueprint-step-1');
    const step2 = document.getElementById('blueprint-step-2');
    const step3 = document.getElementById('blueprint-step-3');

    if (!container || !progressBar) return;

    let ticking = false;

    const updateSteps = () => {
        const rect = container.getBoundingClientRect();
        const vh = window.innerHeight;

        // Start animation when top of the container reaches 80% of viewport height (lower on screen)
        // End animation when it reaches 30% of viewport height (higher on screen)
        const startTrigger = vh * 0.8;
        const endTrigger = vh * 0.3;
        const scrollRange = startTrigger - endTrigger;

        let progress = (startTrigger - rect.top) / scrollRange;

        // Clamp between 0 and 1
        progress = Math.max(0, Math.min(1, progress));

        // Update the horizontal connecting line (desktop)
        progressBar.style.width = `${progress * 100}%`;

        // Sequentially activate the cards based on scroll progress
        if (progress > 0.05) step1?.classList.add('is-active');
        else step1?.classList.remove('is-active');

        if (progress > 0.5) step2?.classList.add('is-active');
        else step2?.classList.remove('is-active');

        if (progress > 0.95) step3?.classList.add('is-active');
        else step3?.classList.remove('is-active');

        ticking = false;
    };

    window.addEventListener('scroll', () => {
        if (!ticking) {
            window.requestAnimationFrame(updateSteps);
            ticking = true;
        }
    }, { passive: true });

    // Initial check on load
    updateSteps();

};
