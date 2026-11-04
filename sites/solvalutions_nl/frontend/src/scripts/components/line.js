const TEETH = { v1: 12, v2: 8 };
const DIRECTION = { v1: 1, v2: -1 };
const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
const desktop = window.matchMedia('(min-width: 64rem)');

const clamp = (value) => Math.min(1, Math.max(0, value));

export const initLine = () => {
    const line = document.querySelector('[data-line]');

    if (!line) return;

    const stations = Array.from(line.querySelectorAll('[data-station]'));
    const belt = line.querySelector('.belt');
    const crate = line.querySelector('.belt-crate');
    const gears = Array.from(line.querySelectorAll('.sc-gr'));
    let frame = null;

    const progress = () => {
        if (reduced) return 1;

        const rect = line.getBoundingClientRect();

        return clamp((window.innerHeight * 0.85 - rect.top) / (window.innerHeight * 0.75));
    };

    const update = (time) => {
        const value = progress();
        const drive = value * 700 + (reduced ? 0 : time / 40);
        const index = Math.min(stations.length - 1, Math.floor(value * stations.length));

        stations.forEach((station, position) => {
            station.classList.toggle('is-active', position === index);
            station.classList.toggle('is-done', position < index);
        });

        gears.forEach((gear) => {
            const angle = Number(gear.dataset.ph) + (DIRECTION[gear.dataset.g] * drive) / TEETH[gear.dataset.g];

            gear.setAttribute('transform', `rotate(${angle.toFixed(2)})`);
        });

        if (belt && crate && desktop.matches) {
            const travel = Math.max(0, belt.clientWidth - crate.offsetWidth);

            crate.style.setProperty('--cx', `${(value * travel).toFixed(1)}px`);
            crate.classList.toggle('is-digital', value > 0.84);
        }
    };

    const compact = window.matchMedia('(max-width: 63.99rem)');
    let lastUpdate = 0;

    const loop = (time) => {
        if (!compact.matches || time - lastUpdate >= 33) {
            lastUpdate = time;
            update(time);
        }

        frame = requestAnimationFrame(loop);
    };

    update(0);

    if (reduced) return;

    const observer = new IntersectionObserver(([entry]) => {
        if (entry.isIntersecting && frame === null) frame = requestAnimationFrame(loop);

        if (!entry.isIntersecting && frame !== null) {
            cancelAnimationFrame(frame);
            frame = null;
        }
    });

    observer.observe(line);
};
