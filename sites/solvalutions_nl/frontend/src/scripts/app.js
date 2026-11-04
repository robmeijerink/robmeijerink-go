import '../styles/app.css';

import { initNavigation } from './components/navigation.js';
import { initDataHighway } from './components/data-highway.js';
import { initEffects } from './components/effects.js';
import { initScene, initSceneFocus } from './components/scene.js';
import { initLine } from './components/line.js';
import { initVignettes } from './components/vignettes.js';
import { initCheck } from './components/check.js';
import { initSecnav } from './components/secnav.js';

document.addEventListener('DOMContentLoaded', () => {
    initNavigation();
    initDataHighway();
    initEffects();
    initScene();
    initSceneFocus();
    initLine();
    initVignettes();
    initCheck();
    initSecnav();
});
