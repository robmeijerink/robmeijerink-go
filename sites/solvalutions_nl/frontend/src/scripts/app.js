import '../styles/app.css'

/**
 * Solvalutions Combined Engine (Kinetic v3.1 + Smart Nav)
 * Bugfix: Engine initialization sequence corrected to prevent NaN coordinates.
 */
document.addEventListener('DOMContentLoaded', () => {
    // --- SMART NAVIGATION CONTROLLER ---
    const nav = document.getElementById('main-nav');
    if (nav) {
        let lastScrollY = window.scrollY;
        let isNavVisible = true;
        let scrollUpAccumulator = 0;

        const handleScroll = () => {
            const currentScrollY = window.scrollY;
            const delta = currentScrollY - lastScrollY;

            // 1. Dynamic Glassmorphism Background
            if (currentScrollY > 50) {
                nav.classList.remove('bg-transparent', 'border-transparent', 'py-2');
                nav.classList.add('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5', 'py-0');
            } else {
                nav.classList.add('bg-transparent', 'border-transparent', 'py-2');
                nav.classList.remove('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5', 'py-0');
            }

            // 2. Hide / Show Logic
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

    // --- KINETIC CIRCUIT ENGINE ---
    const canvas = document.getElementById('flow-canvas');
    if (!canvas) return;
    const ctx = canvas.getContext('2d');

    let width, height, dpr, tracers = [], cardZone = null, pulses = [];
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
        updateCardBounds();
    };

    const updateCardBounds = () => {
        const el = document.getElementById('hero-card-zone');
        if (el) {
            const rect = el.getBoundingClientRect();
            cardZone = { x: rect.left, y: rect.top, w: rect.width, h: rect.height };
        }
    };

    class Pulse {
        constructor(x, y) { this.x = x; this.y = y; this.r = 1; this.opacity = 1; }
        update() { this.r += 3.5; this.opacity -= 0.04; }
        draw() {
            ctx.beginPath(); ctx.arc(this.x, this.y, this.r, 0, Math.PI * 2);
            ctx.strokeStyle = accentColor; ctx.lineWidth = 2; ctx.globalAlpha = Math.max(0, this.opacity); ctx.stroke();
        }
    }

    class Tracer {
        constructor() { this.reset(); }
        reset() {
            const side = Math.floor(Math.random() * 4);
            const speed = 2.5;

            // Spawn points now safely use the initialized 'width' and 'height' variables
            if (side === 0) { this.x = Math.random() * width; this.y = -20; this.vx = 0; this.vy = speed; }
            else if (side === 1) { this.x = width + 20; this.y = Math.random() * height; this.vx = -speed; this.vy = 0; }
            else if (side === 2) { this.x = Math.random() * width; this.y = height + 20; this.vx = 0; this.vy = -speed; }
            else { this.x = -20; this.y = Math.random() * height; this.vx = speed; this.vy = 0; }

            this.history = []; this.maxLength = 25; this.turnTimer = 0;
        }
        update() {
            this.x += this.vx; this.y += this.vy;
            this.turnTimer++;
            if (this.turnTimer > 30 && Math.random() > 0.96) {
                const oldVx = this.vx;
                this.vx = (this.vy === 0) ? 0 : (Math.random() > 0.5 ? 2.5 : -2.5);
                this.vy = (oldVx === 0) ? 0 : (Math.random() > 0.5 ? 2.5 : -2.5);
                this.turnTimer = 0;
            }
            if (cardZone) {
                const m = 10;
                if (this.x > cardZone.x - m && this.x < cardZone.x + cardZone.w + m &&
                    this.y > cardZone.y - m && this.y < cardZone.y + cardZone.h + m) {
                    pulses.push(new Pulse(this.x, this.y));
                    this.reset();
                }
            }
            this.history.push({x: this.x, y: this.y});
            if (this.history.length > this.maxLength) this.history.shift();

            // Out of bounds reset
            if (this.x < -100 || this.x > width + 100 || this.y < -100 || this.y > height + 100) this.reset();
        }
        draw() {
            if (this.history.length < 2) return;

            ctx.beginPath(); ctx.strokeStyle = accentColor; ctx.lineWidth = 1.5; ctx.globalAlpha = 0.35;
            ctx.moveTo(this.history[0].x, this.history[0].y);
            this.history.forEach(p => ctx.lineTo(p.x, p.y)); ctx.stroke();

            ctx.beginPath(); ctx.arc(this.x, this.y, 3, 0, Math.PI * 2);
            ctx.shadowBlur = 8; ctx.shadowColor = accentColor;
            ctx.fillStyle = accentColor; ctx.globalAlpha = 1; ctx.fill();
            ctx.shadowBlur = 0;
        }
    }

    const animate = () => {
        ctx.clearRect(0, 0, width, height);

        // Chip Grid (24px)
        ctx.beginPath(); ctx.strokeStyle = accentColor; ctx.globalAlpha = 0.15; ctx.lineWidth = 0.5;
        for(let i=0; i<width; i+=24) { ctx.moveTo(i, 0); ctx.lineTo(i, height); }
        for(let j=0; j<height; j+=24) { ctx.moveTo(0, j); ctx.lineTo(width, j); }
        ctx.stroke();

        pulses.forEach((p, i) => { p.update(); p.draw(); if (p.opacity <= 0) pulses.splice(i, 1); });
        tracers.forEach(t => { t.update(); t.draw(); });
        requestAnimationFrame(animate);
    };

    const initEngine = () => {
        resize(); // Step 1: Secure screen dimensions first
        for (let i = 0; i < 45; i++) {
            tracers.push(new Tracer()); // Step 2: Spawn tracers with valid coordinates
        }
        animate(); // Step 3: Run loop
    };

    window.addEventListener('resize', resize);
    window.addEventListener('scroll', updateCardBounds, { passive: true });

    // Boot sequence corrected
    initEngine();
});
