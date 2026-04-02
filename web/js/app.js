import { Wizard } from './wizard.js';
import { Results } from './results.js';

const api = {
  async questions() {
    const r = await fetch('/api/questions');
    if (!r.ok) throw new Error(r.status);
    return r.json();
  },
  async profiles() {
    const r = await fetch('/api/profiles');
    if (!r.ok) throw new Error(r.status);
    return r.json();
  },
  async recommend(body) {
    const r = await fetch('/api/recommend', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    });
    if (!r.ok) throw new Error(r.status);
    return r.json();
  },
};

class App {
  constructor() {
    this.views = {
      landing: document.getElementById('view-landing'),
      quiz:    document.getElementById('view-quiz'),
      results: document.getElementById('view-results'),
    };
    this.wizard  = null;
    this.results = null;
  }

  async init() {
    const [questions, profiles] = await Promise.all([
      api.questions(),
      api.profiles(),
    ]);

    this.wizard  = new Wizard(questions, (answers) => this.onQuizDone(answers));
    this.results = new Results();

    this.renderProfiles(profiles);
    this.bind();
    this.restoreTheme();
  }

  show(name) {
    for (const v of Object.values(this.views)) v.classList.remove('active');
    this.views[name].classList.add('active');
    window.scrollTo(0, 0);
  }

  renderProfiles(profiles) {
    const el = document.getElementById('profiles');
    el.innerHTML = profiles.map(p => `
      <button class="profile-card" data-profile="${p.id}">
        <span class="profile-icon">${p.icon}</span>
        <span class="profile-name">${p.name}</span>
        <span class="profile-desc">${p.description}</span>
      </button>
    `).join('');
  }

  bind() {
    document.getElementById('start-quiz').addEventListener('click', () => {
      this.wizard.reset();
      this.show('quiz');
    });

    document.getElementById('profiles').addEventListener('click', async (e) => {
      const card = e.target.closest('[data-profile]');
      if (!card) return;
      card.disabled = true;
      try {
        const data = await api.recommend({ profile: card.dataset.profile });
        this.results.render(data);
        this.show('results');
      } finally {
        card.disabled = false;
      }
    });

    document.getElementById('btn-restart').addEventListener('click', () => {
      this.show('landing');
    });

    document.getElementById('theme-toggle').addEventListener('click', () => {
      const html = document.documentElement;
      const next = html.dataset.theme === 'dark' ? 'light' : 'dark';
      html.dataset.theme = next;
      localStorage.setItem('dp-theme', next);
    });
  }

  restoreTheme() {
    const saved = localStorage.getItem('dp-theme');
    if (saved) document.documentElement.dataset.theme = saved;
  }

  async onQuizDone(answers) {
    const data = await api.recommend({ answers });
    this.results.render(data);
    this.show('results');
  }
}

const app = new App();
app.init().catch((err) => {
  document.getElementById('app').innerHTML = `
    <div style="text-align:center;padding:4rem">
      <p>Failed to load. Is the server running?</p>
      <pre style="color:var(--red)">${err}</pre>
    </div>`;
});
