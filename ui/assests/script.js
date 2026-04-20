// ─── config ────────────────────────────────────────────────────────────────
const BASE = "https://your-backend.com";

async function call(path, body = null) {
  const res = await fetch(`${BASE}${path}`, {
    method: body ? "POST" : "GET",
    headers: { "Content-Type": "application/json" },
    ...(body && { body: JSON.stringify(body) }),
  });
  const text = await res.text();
  if (!res.ok) throw new Error(`HTTP ${res.status}: ${text}`);
  return text;
}

// ─── DOM refs ──────────────────────────────────────────────────────────────
const urlInput = document.getElementById("url-input");
const forgeBtn = document.getElementById("btn-forge");
const shortenBtn = document.getElementById("btn-shorten");
const codeOutput = document.getElementById("code-output");
const copyBtn = document.getElementById("btn-copy");

const statTotal = document.getElementById("stat-total-links");
const statClicks = document.getElementById("stat-total-clicks");
const statFlagged = document.getElementById("stat-flagged");
const statAvg = document.getElementById("stat-avg-redirect");

// ─── validation ────────────────────────────────────────────────────────────
function getURL() {
  const val = urlInput.value.trim();
  if (!val) {
    urlInput.focus();
    return null;
  }
  try {
    new URL(val);
    return val;
  } catch {
    urlInput.setCustomValidity("Enter a valid URL");
    urlInput.reportValidity();
    return null;
  }
}

urlInput.addEventListener("input", () => urlInput.setCustomValidity(""));

// ─── forge ─────────────────────────────────────────────────────────────────
forgeBtn.addEventListener("click", async () => {
  const url = getURL();
  if (!url) return;

  setLoading(forgeBtn, true);
  try {
    const result = await call("/forge", { url });
    showCode(result);
  } catch (e) {
    showCode(`Error: ${e.message}`);
  } finally {
    setLoading(forgeBtn, false);
  }
});

// ─── shorten ───────────────────────────────────────────────────────────────
shortenBtn.addEventListener("click", async () => {
  const url = getURL();
  if (!url) return;

  setLoading(shortenBtn, true);
  try {
    const result = await call("/shorten", { url });
    showCode(result);
  } catch (e) {
    showCode(`Error: ${e.message}`);
  } finally {
    setLoading(shortenBtn, false);
  }
});

// ─── copy ──────────────────────────────────────────────────────────────────
copyBtn.addEventListener("click", () => {
  navigator.clipboard.writeText(codeOutput.textContent).then(() => {
    copyBtn.textContent = "Copied!";
    setTimeout(() => (copyBtn.textContent = "Copy"), 1500);
  });
});

// ─── stats poll ────────────────────────────────────────────────────────────
async function fetchStats() {
  try {
    const raw = await call("/stats");
    const data = JSON.parse(raw);
    statTotal.textContent = data.totalLinks ?? "—";
    statClicks.textContent = data.totalClicks ?? "—";
    statFlagged.textContent = data.flagged ?? "—";
    statAvg.textContent =
      data.avgRedirectMs != null ? `${Math.round(data.avgRedirectMs)} ms` : "—";
  } catch (e) {
    console.error("[stats poll]", e);
  }
}

fetchStats();
setInterval(fetchStats, 5000);

// ─── helpers ───────────────────────────────────────────────────────────────
function showCode(text) {
  codeOutput.textContent = text;
  codeOutput.closest("[data-output]").style.display = "block";
}

function setLoading(btn, on) {
  btn.disabled = on;
  btn.dataset.label ??= btn.textContent;
  btn.textContent = on ? "…" : btn.dataset.label;
}
