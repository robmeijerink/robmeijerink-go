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

    const loop = (time) => {
        draw(time);
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

    const set = (value) => {
        if (value) stage.setAttribute('data-focus', value);
        else stage.removeAttribute('data-focus');
    };

    stage.querySelectorAll('[data-focus]').forEach((callout) => {
        const value = callout.getAttribute('data-focus');

        callout.addEventListener('pointerenter', () => set(value));
        callout.addEventListener('pointerleave', () => set(null));
        callout.addEventListener('focus', () => set(value));
        callout.addEventListener('blur', () => set(null));
    });
};
