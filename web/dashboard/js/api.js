const API_BASE = "http://localhost:8080";

async function apiGet(path) {
    const res = await fetch(API_BASE + path);
    return res.json();
}

async function getSummary() {
    return apiGet("/analysis/summary");
}

async function getStats() {
    return apiGet("/stats");
}

async function getTrends() {
    return apiGet("/trends");
}

async function getRecords() {
    return apiGet("/records");
}

async function getPatterns() {
    return apiGet("/analysis/patterns");
}
