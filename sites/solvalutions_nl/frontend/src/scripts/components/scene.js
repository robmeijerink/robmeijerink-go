const TEETH = { g1: 14, g2: 10, g3: 8 };
const DIRECTION = { g1: 1, g2: -1, g3: -1 };
const BELT_LENGTH = 6.2;
const COS = 0.8660254;
const SIN = 0.5;
const UNIT = 34;

export const initScene = () => {
    const stage = document.querySelector('[data-scene]');

    if (!stage) return;

    const gears = Array.from(stage.querySelectorAll('.sc-gr'));
    const crates = Array.from(stage.querySelectorAll('.sc-crate'));
    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    let frame = null;

    const draw = (time) => {
        const seconds = reduced ? 0 : time / 1000;
        const drive = seconds * 150 + window.scrollY * 2.4;

        gears.forEach((gear) => {
            const angle = Number(gear.dataset.ph) + (DIRECTION[gear.dataset.g] * drive) / TEETH[gear.dataset.g];

            gear.setAttribute('transform', `rotate(${angle.toFixed(2)})`);
        });

        crates.forEach((crate) => {
            const travel = (Number(crate.dataset.u) + seconds * 0.05 + window.scrollY * 0.0004) % 1;
            const x = travel * BELT_LENGTH * COS * UNIT;
            const y = travel * BELT_LENGTH * SIN * UNIT;

            crate.setAttribute('transform', `translate(${x.toFixed(1)} ${y.toFixed(1)})`);
        });
    };

    const compact = window.matchMedia('(max-width: 63.99rem)');
    let lastDraw = 0;

    const loop = (time) => {
        if (!compact.matches || time - lastDraw >= 33) {
            lastDraw = time;
            draw(time);
        }

        frame = requestAnimationFrame(loop);
    };

    draw(0);

    if (reduced) return;

    const observer = new IntersectionObserver(([entry]) => {
        if (entry.isIntersecting && frame === null) frame = requestAnimationFrame(loop);

        if (!entry.isIntersecting && frame !== null) {
            cancelAnimationFrame(frame);
            frame = null;
        }
    });

    observer.observe(stage);
};

export const initSceneFocus = () => {
    const stage = document.querySelector('[data-scene]');

    if (!stage) return;

    const pins = Array.from(stage.querySelectorAll('.sc-dot'));
    const first = pins[0];
    const desktop = window.matchMedia('(min-width: 64rem)');
    let touched = false;

    const set = (value) => {
        if (value) stage.setAttribute('data-focus', value);
        else stage.removeAttribute('data-focus');
    };

    const close = () => pins.forEach((pin) => pin.classList.remove('is-open'));

    const settle = () => {
        if (touched || !desktop.matches || !first) return;

        first.classList.add('is-open');
        set(first.getAttribute('data-focus'));
    };

    pins.forEach((pin) => {
        const value = pin.getAttribute('data-focus');

        pin.addEventListener('pointerenter', () => {
            touched = true;
            close();
            set(value);
        });
        pin.addEventListener('pointerleave', () => set(null));
        pin.addEventListener('focus', () => {
            touched = true;
            close();
            set(value);
        });
        pin.addEventListener('blur', () => set(null));
        pin.addEventListener('click', () => {
            touched = true;
            close();
            pin.classList.add('is-open');
            set(value);
        });
    });

    document.addEventListener('click', (event) => {
        if (!event.target.closest('.sc-dot')) {
            close();
            set(null);
        }
    });

    document.querySelectorAll('#hero-section .pil').forEach((item) => {
        const value = item.getAttribute('data-focus');

        item.addEventListener('pointerenter', () => set(value));
        item.addEventListener('pointerleave', () => set(null));
    });

    if (!window.matchMedia('(prefers-reduced-motion: reduce)').matches) setTimeout(settle, 1200);
    else settle();
};
