export class Wizard {
  constructor(questions, onComplete) {
    this.questions  = questions;
    this.onComplete = onComplete;
    this.step       = 0;
    this.answers    = {};

    this.container   = document.getElementById('question-container');
    this.progressBar = document.getElementById('progress-bar');
    this.progressTxt = document.getElementById('progress-text');
    this.btnBack     = document.getElementById('btn-back');
    this.btnNext     = document.getElementById('btn-next');

    this.btnBack.addEventListener('click', () => this.prev());
    this.btnNext.addEventListener('click', () => this.next());
  }

  reset() {
    this.step = 0;
    this.answers = {};
    this.render();
  }

  render() {
    const q     = this.questions[this.step];
    const total = this.questions.length;
    const sel   = this.answers[q.id] || [];

    this.progressBar.style.width = `${((this.step + 1) / total) * 100}%`;
    this.progressTxt.textContent = `${this.step + 1} / ${total}`;
    this.btnBack.disabled        = this.step === 0;
    this.btnNext.textContent     = this.step === total - 1 ? 'See Results →' : 'Next →';

    this.container.innerHTML = `
      <div class="question" role="group" aria-labelledby="q-title">
        <h2 id="q-title" class="question-title">${q.text}</h2>
        ${q.subtitle ? `<p class="question-subtitle">${q.subtitle}</p>` : ''}
        <div class="options-grid" role="${q.multi_select ? 'group' : 'radiogroup'}">
          ${q.options.map(o => `
            <button class="option-card ${sel.includes(o.id) ? 'selected' : ''}"
                    data-id="${o.id}"
                    role="${q.multi_select ? 'checkbox' : 'radio'}"
                    aria-checked="${sel.includes(o.id)}"
                    type="button">
              ${o.icon ? `<span class="option-icon">${o.icon}</span>` : ''}
              <span class="option-label">${o.label}</span>
              ${o.desc ? `<span class="option-desc">${o.desc}</span>` : ''}
            </button>
          `).join('')}
        </div>
        ${q.multi_select ? '<p class="hint">Select all that apply</p>' : ''}
      </div>`;

    this.container.querySelectorAll('.option-card').forEach(card => {
      card.addEventListener('click', () => this.select(q, card.dataset.id));
    });
  }

  select(question, optionId) {
    if (question.multi_select) {
      const cur = this.answers[question.id] || [];
      const idx = cur.indexOf(optionId);
      if (idx >= 0) cur.splice(idx, 1);
      else cur.push(optionId);
      this.answers[question.id] = cur;
    } else {
      this.answers[question.id] = [optionId];
    }
    this.render();
  }

  prev() {
    if (this.step > 0) {
      this.step--;
      this.render();
    }
  }

  next() {
    const q   = this.questions[this.step];
    const sel = this.answers[q.id];

    if (!sel || sel.length === 0) {
      const grid = this.container.querySelector('.options-grid');
      grid.classList.remove('shake');
      void grid.offsetWidth; // reflow to restart animation
      grid.classList.add('shake');
      return;
    }

    if (this.step === this.questions.length - 1) {
      const formatted = Object.entries(this.answers).map(([qid, choices]) => ({
        question_id: qid,
        choices,
      }));
      this.onComplete(formatted);
    } else {
      this.step++;
      this.render();
    }
  }
}
