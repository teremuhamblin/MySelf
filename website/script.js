console.log("MySelf website loaded.");
// Inject layout (header + footer)
document.getElementById("layout").innerHTML = `
<header>
    <img src="assets/logo.svg" alt="Logo" class="logo">
    <nav>
        <a href="index.html">Accueil</a>
        <a href="pages/identity.html">Identité</a>
        <a href="pages/doctrine.html">Doctrine</a>
        <a href="pages/routines.html">Routines</a>
        <a href="pages/goals.html">Objectifs</a>
        <a href="pages/tracking.html">Tracking</a>
        <a href="pages/dashboard.html">Dashboard</a>
    </nav>
    <button class="theme-toggle" onclick="toggleTheme()">Mode</button>
</header>
`;

// Dark mode
function toggleTheme() {
    const current = document.documentElement.getAttribute("data-theme");
    document.documentElement.setAttribute(
        "data-theme",
        current === "dark" ? "light" : "dark"
    );
}

// Dashboard example logic
if (window.location.pathname.includes("dashboard.html")) {
    console.log("Dashboard loaded");
}
