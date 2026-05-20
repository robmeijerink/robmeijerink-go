import '../styles/app.css';

import { initNavigation } from './components/navigation.js';
import { initDataHighway, initBlueprintSteps } from './pages/home.js';

document.addEventListener('DOMContentLoaded', () => {
    initNavigation();
    initDataHighway();
    initBlueprintSteps();
});
