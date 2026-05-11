import '../styles/app.css'

/**
 * Solvalutions Combined Engine (Data Highway v6.5)
 * Features: Slower, stable movement. Bouncing collisions (no disappearing). Solid state.
 */
document.addEventListener('DOMContentLoaded', () => {

    const nav = document.getElementById('main-nav');
    if (nav) {
        let lastScrollY = window.scrollY;
        let isNavVisible = true;
        let scrollUpAccumulator = 0;

        const handleScroll = () => {
            const currentScrollY = window.scrollY;
            const delta = currentScrollY - lastScrollY;

            if (currentScrollY > 50) {
                nav.classList.remove('bg-transparent', 'border-transparent', 'py-2');
                nav.classList.add('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5', 'py-0');
            } else {
                nav.classList.add('bg-transparent', 'border-transparent', 'py-2');
                nav.classList.remove('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5', 'py-0');
            }

            if (currentScrollY <= 150) {
                if (!isNavVisible) {
                    nav.style.transform = 'translateY(0)';
                    isNavVisible = true;
                }
                scrollUpAccumulator = 0;
            } else if (delta > 0) {
                if (isNavVisible) {
                    nav.style.transform = 'translateY(-100%)';
                    isNavVisible = false;
                }
                scrollUpAccumulator = 0;
            } else if (delta < 0) {
                scrollUpAccumulator += Math.abs(delta);
                if (!isNavVisible && scrollUpAccumulator > 80) {
                    nav.style.transform = 'translateY(0)';
                    isNavVisible = true;
                }
            }
            lastScrollY = currentScrollY;
        };

        window.addEventListener('scroll', handleScroll, { passive: true });
    }

    const canvas = document.getElementById('flow-canvas');
    if (!canvas) return;
    const ctx = canvas.getContext('2d');

    let width, height, dpr, tracers = [], pulses = [];
    let highwayMinY = 0;
    let highwayMaxY = 0;
    const accentColor = '#B84A2B';

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
        const heroContent = document.getElementById('hero-content');
        const casesHeader = document.getElementById('cases-header');

        if (heroContent && casesHeader) {
            const heroRect = heroContent.getBoundingClientRect();
            const casesRect = casesHeader.getBoundingClientRect();

            const heroBottom = heroRect.bottom + window.scrollY;
            const casesTop = casesRect.top + window.scrollY;

            highwayMinY = heroBottom + 20;
            highwayMaxY = casesTop - 20;
        }
    };

    class Pulse {
        constructor(x, y) {
            this.x = x;
            this.y = y;
            this.r = 1;
            this.opacity = 0.6;
        }
        update() {
            this.r += 1.2;
            this.opacity -= 0.03;
        }
        draw() {
            const renderY = this.y - window.scrollY;
            if (renderY < -50 || renderY > height + 50) return;

            ctx.beginPath();
            ctx.arc(this.x, renderY, this.r, 0, Math.PI * 2);
            ctx.strokeStyle = accentColor;
            ctx.lineWidth = 1.5;
            ctx.globalAlpha = Math.max(0, this.opacity);

            ctx.shadowBlur = 12;
            ctx.shadowColor = accentColor;

            ctx.stroke();
            ctx.shadowBlur = 0;
        }
    }

    class Tracer {
        constructor() { this.reset(); }
        reset() {
            const startLeft = Math.random() > 0.5;
            const speed = 1.2; // Reduced speed by roughly 50% for smooth, calm flow

            this.y = Math.random() * (highwayMaxY - highwayMinY) + highwayMinY;

            if (startLeft) {
                this.x = -20;
                this.mainVx = speed;
            } else {
                this.x = width + 20;
                this.mainVx = -speed;
            }

            this.vx = this.mainVx;
            this.vy = 0;
            this.history = [];
            this.maxLength = 40; // Longer tail to accommodate slower speed seamlessly
            this.turnTimer = 0;
            this.collisionCooldown = 0; // Prevents multiple pulses during a single collision
        }

        update() {
            this.x += this.vx;
            this.y += this.vy;
            this.turnTimer++;

            if (this.collisionCooldown > 0) {
                this.collisionCooldown--;
            }

            // Smoother, less erratic turning interval
            if (this.turnTimer > 25 && Math.random() > 0.60) {
                if (this.vy === 0) {
                    this.vx = 0;
                    this.vy = (Math.random() > 0.5 ? 1.2 : -1.2);
                } else {
                    this.vy = 0;
                    this.vx = this.mainVx;
                }
                this.turnTimer = 0;
            }

            if (this.y <= highwayMinY) {
                this.y = highwayMinY;
                this.vy = 0;
                this.vx = this.mainVx;
            } else if (this.y >= highwayMaxY) {
                if (this.vy > 0) {
                    pulses.push(new Pulse(this.x, this.y));
                }
                this.y = highwayMaxY;
                this.vy = 0;
                this.vx = this.mainVx;
            }

            this.history.push({x: this.x, y: this.y});
            if (this.history.length > this.maxLength) this.history.shift();

            if (this.x < -100 || this.x > width + 100) this.reset();
        }

        draw() {
            if (this.history.length < 2) return;

            const renderY = this.y - window.scrollY;
            if (renderY < -100 || renderY > height + 100) return;

            ctx.beginPath();
            ctx.strokeStyle = accentColor;
            ctx.lineWidth = 1.5;
            ctx.globalAlpha = 0.35;
            ctx.moveTo(this.history[0].x, this.history[0].y - window.scrollY);
            this.history.forEach(p => ctx.lineTo(p.x, p.y - window.scrollY));
            ctx.stroke();

            // Head without shadowBlur ensures solid, non-flickering drawing
            ctx.beginPath();
            ctx.arc(this.x, renderY, 3, 0, Math.PI * 2);
            ctx.fillStyle = accentColor;
            ctx.globalAlpha = 1;
            ctx.fill();
        }
    }

    const animate = () => {
        ctx.clearRect(0, 0, width, height);

        const scrollOffset = window.scrollY % 24;

        ctx.beginPath();
        ctx.strokeStyle = accentColor;
        ctx.globalAlpha = 0.15;
        ctx.lineWidth = 0.5;
        for(let i=0; i<width; i+=24) { ctx.moveTo(i, 0); ctx.lineTo(i, height); }
        for(let j= -scrollOffset; j<height; j+=24) { ctx.moveTo(0, j); ctx.lineTo(width, j); }
        ctx.stroke();

        for (let i = 0; i < tracers.length; i++) {
            for (let j = i + 1; j < tracers.length; j++) {
                const dx = tracers[i].x - tracers[j].x;
                const dy = tracers[i].y - tracers[j].y;

                if (dx * dx + dy * dy < 150) {
                    if (tracers[i].collisionCooldown === 0 && tracers[j].collisionCooldown === 0) {
                        pulses.push(new Pulse((tracers[i].x + tracers[j].x) / 2, (tracers[i].y + tracers[j].y) / 2));

                        // Handshake: Bounce off vertically instead of disappearing
                        tracers[i].vy = tracers[i].vy === 0 ? (Math.random() > 0.5 ? 1.2 : -1.2) : tracers[i].vy * -1;
                        tracers[j].vy = tracers[j].vy === 0 ? (Math.random() > 0.5 ? 1.2 : -1.2) : tracers[j].vy * -1;

                        // Make sure they resume horizontal movement quickly
                        tracers[i].vx = 0;
                        tracers[j].vx = 0;

                        tracers[i].turnTimer = 0;
                        tracers[j].turnTimer = 0;

                        // Prevent rapid multi-collisions
                        tracers[i].collisionCooldown = 30;
                        tracers[j].collisionCooldown = 30;
                    }
                }
            }
        }

        pulses.forEach((p, i) => {
            p.update();
            p.draw();
            if (p.opacity <= 0) pulses.splice(i, 1);
        });

        tracers.forEach(t => { t.update(); t.draw(); });
        requestAnimationFrame(animate);
    };

    const initEngine = () => {
        resize();
        for (let i = 0; i < 8; i++) {
            tracers.push(new Tracer());
        }
        animate();
    };

    window.addEventListener('resize', resize);
    setTimeout(updateHighwayBounds, 500);

    initEngine();
});
