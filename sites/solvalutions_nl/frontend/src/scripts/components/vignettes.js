const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

export const initVignettes = () => {
    const gears = Array.from(document.querySelectorAll('.sc-gr[data-vd]'));

    if (!gears.length || reduced) return;

    const visible = new Set();
    let frame = null;

    let lastTick = 0;

    const loop = (time) => {
        if (time - lastTick < 33) {
            frame = requestAnimationFrame(loop);
            return;
        }

        lastTick = time;

        visible.forEach((gear) => {
            const angle = Number(gear.dataset.ph) + (Number(gear.dataset.vd) * time) / (28 + Number(gear.dataset.n) * 2);

            gear.setAttribute('transform', `rotate(${angle.toFixed(2)})`);
        });

        frame = visible.size ? requestAnimationFrame(loop) : null;
    };

    const observer = new IntersectionObserver((entries) => {
        entries.forEach((entry) => {
            const gear = entry.target;

            if (entry.isIntersecting) visible.add(gear);
            else visible.delete(gear);
        });

        if (visible.size && frame === null) frame = requestAnimationFrame(loop);
    });

    gears.forEach((gear) => observer.observe(gear));
};
