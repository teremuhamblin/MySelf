async function loadDashboard() {
    const summary = await getSummary();
    const stats = await getStats();
    const trends = await getTrends();
    const records = await getRecords();

    // Résumé
    document.getElementById("summary-text").innerText = summary.text;

    // Graphiques
    renderTypesChart(document.getElementById("typesChart"), stats);
    renderTagsChart(document.getElementById("tagsChart"), stats);
    renderTrendsChart(document.getElementById("trendsChart"), trends);

    // Liste des records
    const list = document.getElementById("records-list");
    list.innerHTML = "";

    records.slice(-10).reverse().forEach(r => {
        const li = document.createElement("li");
        li.innerHTML = `<strong>${r.type}</strong> — ${r.content.substring(0, 60)}...`;
        list.appendChild(li);
    });
}

loadDashboard();
