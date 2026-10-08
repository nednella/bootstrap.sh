---
name: unslop
description: Remove AI writing patterns from prose that is written into files or posted outside the session, such as PR and issue text, commit messages, docs and code comments. Use it before writing or posting that prose, or when asked to "unslop", "humanize" or "de-AI" text; not for chat replies.
---

# Unslop

Edit prose to remove AI patterns and leave it plain.

Scope: text that lands in a file or outside this session. That means PR and issue text, commit messages, docs and code comments. Chat replies are out of scope.

## Process

1. Scan for the patterns below.
2. Rewrite. Keep the meaning and the intended tone.
3. Check the result reads as plain English:
   - Prefer short words.
   - Use the active voice.
   - Put one idea in each sentence.
   - Use no figures of speech.
   - Use no jargon where an everyday word works.
4. Ask "what makes this read as AI-written?" Fix what remains.

## Patterns to find and fix

### Content

1. **Inflated significance**: "pivotal moment", "testament to", "evolving landscape", "setting the stage for", "indelible mark", "deeply rooted". Cut the puffery and state what happened.
2. **Name-dropping**: a list of outlets or names with no context. Pick one and say what it said.
3. **Shallow -ing phrases**: "highlighting...", "ensuring...", "reflecting...", "showcasing...", "fostering...". Delete them, or expand with a real source.
4. **Promotional language**: "nestled", "vibrant", "breathtaking", "groundbreaking", "renowned", "stunning", "must-visit". Describe it plainly.
5. **Vague attributions**: "Experts believe", "Industry reports suggest", "Some critics argue". Name the source or delete the claim.
6. **Formulaic challenges**: "Despite challenges... continues to thrive." Replace with specific facts.

### Language

7. **AI vocabulary**: additionally, crucial, delve, enduring, enhance, fostering, garner, interplay, intricate, landscape (abstract), pivotal, showcase, tapestry (abstract), testament, underscore, vibrant. Use plain words.
8. **Copula avoidance**: "serves as", "stands as", "boasts", "features". Say "is" or "has".
9. **Negative parallelisms**: "It's not just X, it's Y." State the point.
10. **Rule of three**: ideas forced into groups of three. Use the real number.
11. **Synonym cycling**: protagonist, main character, central figure and hero in one paragraph. Pick one word and repeat it.
12. **False ranges**: "from X to Y" where X and Y are not on a real scale. List the topics.

### Style

13. **Em dashes**: do not use them at all. Use periods, commas or parentheses. Em dashes are an AI tell.
14. **Colon overuse**: a colon is fine before a list or an example, not as a link mid-sentence. "If you're coming from traditional automation: instead of registering event handlers, you describe conditions" gains nothing from the colon. Drop the comparison and let the point stand: "Describe when the scheduler should fire in plain English."
15. **Too much bold**: do not bold every proper noun or acronym.
16. **Inline-header lists**: "**Performance:** Performance improved..." Turn them into prose.
17. **Title case headings**: use sentence case.
18. **Decorative emojis**: remove them from headings and bullets.
19. **Curly quotes**: use straight quotes.

### Chatbot leftovers

20. **Chatbot phrases**: "I hope this helps!", "Let me know if...", "Of course!", "Certainly!" Remove them.
21. **Cutoff disclaimers**: "While specific details are limited..." Find the source or remove the line.
22. **Sycophantic tone**: "Great question! You're absolutely right!" Get to the point.

### Filler

23. **Filler phrases**: "In order to" becomes "To". "Due to the fact that" becomes "Because". "It is important to note that" goes.
24. **Heavy hedging**: "could potentially possibly be argued that it might" becomes "may".
25. **Generic endings**: "The future looks bright." State a specific plan or fact.

Adapted from poteto/noodle (MIT).
