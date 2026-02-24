(function () {
  'use strict';

  // Bail if user prefers reduced motion
  if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
    return;
  }

  var observer = new IntersectionObserver(
    function (entries) {
      entries.forEach(function (entry) {
        if (!entry.isIntersecting) return;

        entry.target.classList.add('visible');

        // Stagger child trip-cards
        var cards = entry.target.querySelectorAll('.trip-card');
        cards.forEach(function (card, i) {
          card.style.transitionDelay = (i * 0.08) + 's';
        });

        observer.unobserve(entry.target);
      });
    },
    { threshold: 0.1 }
  );

  document.querySelectorAll('.reveal').forEach(function (el) {
    observer.observe(el);
  });
})();
