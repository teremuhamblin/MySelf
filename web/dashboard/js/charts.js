function renderTypesChart(ctx, stats) {
    new Chart(ctx, {
        type: "pie",
        data: {
            labels: Object.keys(stats.by_type),
            datasets: [{
                data: Object.values(stats.by_type),
                backgroundColor: ["#4CAF50", "#2196F3", "#FFC107", "#E91E63"]
            }]
        }
    });
}

function renderTagsChart(ctx, stats) {
    new Chart(ctx, {
        type: "bar",
        data: {
            labels: Object.keys(stats.tags),
            datasets: [{
                label: "Tags",
                data: Object.values(stats.tags),
                backgroundColor: "#2196F3"
            }]
        }
    });
}

function renderTrendsChart(ctx, trends) {
    new Chart(ctx, {
        type: "line",
        data: {
            labels: Object.keys(trends),
            datasets: [{
                label: "Entrées par mois",
                data: Object.values(trends),
                borderColor: "#E91E63",
                fill: false
            }]
        }
    });
}
