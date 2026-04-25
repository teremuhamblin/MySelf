console.log("MySelf website loaded.");

//
// ===============================
// 1. INJECTION DU HEADER + FOOTER
// ===============================
//

function injectLayout() {
    const layout = document.getElementById("layout");

    layout.innerHTML = `
        <header>
            <img src="assets/logo.svg" alt="Logo" class="logo">

            <nav class="main-nav">
                <a href="index.html" data-page="home">Accueil</a>
                <a href="pages/identity.html" data-page="identity">Identité</a>
                <a href="pages/doctrine.html" data-page="doctrine">Doctrine</a>
                <a href="pages/routines.html" data-page="routines">Routines</a>
                <a href="pages/goals.html" data-page="goals">Objectifs</a>
                <a href="pages/tracking.html" data-page="tracking">Tracking</a>
                <a href="pages/dashboard.html" data-page="dashboard">Dashboard</a>
            </nav>

            <button class="theme-toggle" id="themeToggle">Mode</button>
        </header>

        <footer class="site-footer">
            <p>© 2026 – MySelf Project</p>
        </footer>
    `;
}

injectLayout();


//
// ===============================
// 2. DARK MODE
// ===============================
//

function toggleTheme() {
    const current = document.documentElement.getAttribute("data-theme");
    const next = current === "dark" ? "light" : "dark";
    document.documentElement.setAttribute("data-theme", next);
    localStorage.setItem("theme", next);
}

document.getElementById("themeToggle").addEventListener("click", toggleTheme);

// Charger le thème sauvegardé
const savedTheme = localStorage.getItem("theme");
if (savedTheme) document.documentElement.setAttribute("data-theme", savedTheme);


//
// ===============================
// 3. NAVIGATION ACTIVE
// ===============================
//

function updateActiveLink() {
    const links = document.querySelectorAll("nav a");
    const path = window.location.pathname;

    links.forEach(link => {
        link.classList.toggle("active", path.includes(link.getAttribute("href")));
    });
}

updateActiveLink();


//
// ===============================
// 4. SPA — NAVIGATION DYNAMIQUE
// ===============================
//

async function loadPage(url) {
    const main = document.querySelector("main");

    main.innerHTML = `<p>Chargement...</p>`;

    try {
        const response = await fetch(url);
        const html = await response.text();
        main.innerHTML = html;
    } catch (e) {
        main.innerHTML = `<p>Erreur lors du chargement.</p>`;
    }
}

document.addEventListener("click", (e) => {
    const link = e.target.closest("a[data-page]");
    if (!link) return;

    e.preventDefault();
    const url = link.getAttribute("href");

    history.pushState({}, "", url);
    loadPage(url);
    updateActiveLink();
});

// Navigation via boutons du navigateur
window.addEventListener("popstate", () => {
    loadPage(window.location.pathname);
    updateActiveLink();
});
