/* ─── CANVAS BACKGROUND ─── */
const canvas = document.getElementById('bg-canvas');
const ctx = canvas.getContext('2d');
let W, H, nodes = [], frame = 0;

function resize() {
  W = canvas.width = window.innerWidth;
  H = canvas.height = window.innerHeight;
}

function initNodes() {
  nodes = [];
  const count = Math.floor((W * H) / 18000);
  for (let i = 0; i < count; i++) {
    nodes.push({
      x: Math.random() * W,
      y: Math.random() * H,
      vx: (Math.random() - 0.5) * 0.25,
      vy: (Math.random() - 0.5) * 0.25,
      r: Math.random() * 1.5 + 0.5,
      opacity: Math.random() * 0.4 + 0.1
    });
  }
}

function drawBg() {
  ctx.clearRect(0, 0, W, H);

  // Draw connections
  for (let i = 0; i < nodes.length; i++) {
    for (let j = i + 1; j < nodes.length; j++) {
      const dx = nodes[j].x - nodes[i].x;
      const dy = nodes[j].y - nodes[i].y;
      const dist = Math.sqrt(dx * dx + dy * dy);
      if (dist < 130) {
        const alpha = (1 - dist / 130) * 0.12;
        ctx.beginPath();
        ctx.moveTo(nodes[i].x, nodes[i].y);
        ctx.lineTo(nodes[j].x, nodes[j].y);
        ctx.strokeStyle = `rgba(255,255,255,${alpha})`;
        ctx.lineWidth = 0.5;
        ctx.stroke();
      }
    }
  }

  // Draw nodes
  nodes.forEach(n => {
    ctx.beginPath();
    ctx.arc(n.x, n.y, n.r, 0, Math.PI * 2);
    ctx.fillStyle = `rgba(255,255,255,${n.opacity})`;
    ctx.fill();

    n.x += n.vx;
    n.y += n.vy;
    if (n.x < 0 || n.x > W) n.vx *= -1;
    if (n.y < 0 || n.y > H) n.vy *= -1;
  });

  frame++;
  requestAnimationFrame(drawBg);
}

resize();
initNodes();
drawBg();
window.addEventListener('resize', () => { resize(); initNodes(); });


/* ─── URL SHORTENER LOGIC ─── */
const urlInput = document.getElementById('url-input');
const shortenBtn = document.getElementById('shorten-btn');
const resultPanel = document.getElementById('result-panel');
const resultUrlEl = document.getElementById('result-url');
const copyBtn = document.getElementById('copy-btn');
const qrBtn = document.getElementById('qr-btn');
const newBtn = document.getElementById('new-btn');

// Mock short URL generator
function generateShortCode(len = 5) {
  const chars = 'abcdefghijklmnopqrstuvwxyz0123456789';
  return Array.from({ length: len }, () => chars[Math.floor(Math.random() * chars.length)]).join('');
}

function isValidUrl(str) {
  try {
    new URL(str.startsWith('http') ? str : 'https://' + str);
    return true;
  } catch { return false; }
}

// Animate counter
function animateCount(el, target, duration = 800) {
  const start = parseInt(el.textContent.replace(/,/g, '')) || 0;
  const startTime = performance.now();
  const format = n => n.toLocaleString();
  const step = (now) => {
    const progress = Math.min((now - startTime) / duration, 1);
    const ease = 1 - Math.pow(1 - progress, 3);
    el.textContent = format(Math.round(start + (target - start) * ease));
    if (progress < 1) requestAnimationFrame(step);
  };
  requestAnimationFrame(step);
}

let totalLinks = 12847;
let totalClicks = 3249831;

function updateStats() {
  animateCount(document.getElementById('stat-links'), totalLinks);
  animateCount(document.getElementById('stat-clicks'), totalClicks);
}

setTimeout(updateStats, 600);

// Shorten
shortenBtn.addEventListener('click', () => shorten());
urlInput.addEventListener('keydown', e => { if (e.key === 'Enter') shorten(); });

let currentShortUrl = '';

async function shorten() {
  const raw = urlInput.value.trim();
  if (!raw) { shake(urlInput); return; }

  const url = raw.startsWith('http') ? raw : 'https://' + raw;
  if (!isValidUrl(url)) { shake(urlInput); showToast('⚠ Invalid URL'); return; }

  // Loading state
  shortenBtn.classList.add('loading');
  shortenBtn.disabled = true;

  // Simulate API delay
  await new Promise(r => setTimeout(r, 700 + Math.random() * 400));

  const code = generateShortCode();
  currentShortUrl = `frge.io/${code}`;

  resultUrlEl.href = `https://${currentShortUrl}`;
  resultUrlEl.textContent = currentShortUrl;
  resultPanel.classList.add('visible');

  shortenBtn.classList.remove('loading');
  shortenBtn.disabled = false;

  // Update stats
  totalLinks++;
  totalClicks += Math.floor(Math.random() * 15 + 3);
  updateStats();
}

function shake(el) {
  el.style.animation = 'none';
  el.offsetHeight;
  el.style.animation = 'shake 0.4s ease';
  setTimeout(() => el.style.animation = '', 400);
}

// Copy
copyBtn.addEventListener('click', async () => {
  try {
    await navigator.clipboard.writeText(`https://${currentShortUrl}`);
    copyBtn.classList.add('copied');
    copyBtn.innerHTML = `<svg width="14" height="14" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"/></svg>`;
    showToast('✓  Copied to clipboard');
    setTimeout(() => {
      copyBtn.classList.remove('copied');
      copyBtn.innerHTML = `<svg width="14" height="14" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"/><path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/></svg>`;
    }, 2000);
  } catch { showToast('Copy failed'); }
});

// New / reset
newBtn.addEventListener('click', () => {
  resultPanel.classList.remove('visible');
  urlInput.value = '';
  urlInput.focus();
});

// QR button
qrBtn.addEventListener('click', () => {
  const qrUrl = `https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=${encodeURIComponent('https://' + currentShortUrl)}&bgcolor=0a0a0a&color=f0f0f0`;
  window.open(qrUrl, '_blank');
});

// Toast
function showToast(msg) {
  const toast = document.getElementById('toast');
  toast.textContent = msg;
  toast.classList.add('show');
  setTimeout(() => toast.classList.remove('show'), 2500);
}

// Shake keyframe injection
const style = document.createElement('style');
style.textContent = `
@keyframes shake {
  0%,100% { transform: translateX(0); }
  20% { transform: translateX(-6px); }
  40% { transform: translateX(6px); }
  60% { transform: translateX(-4px); }
  80% { transform: translateX(4px); }
}`;
document.head.appendChild(style);

// GitHub star count (mock update after load)
setTimeout(() => {
  const count = document.getElementById('gh-count');
  if (count) animateCount(count, 248, 1200);
}, 1000);
