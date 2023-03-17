const html = document.querySelector('.html textarea');
const css = document.querySelector('.css textarea');
const js = document.querySelector('.js textarea');
const preview = document.querySelector('.preview iframe');

function updatePreview() {
  const htmlContent = html.value;
  const cssContent = `<style>${css.value}</style>`;
  const jsContent = `<script>${js.value}</script>`;
  preview.contentDocument.write(`${htmlContent}${cssContent}${jsContent}`);
}

html.addEventListener('input', updatePreview);
css.addEventListener('input', updatePreview);
js.addEventListener('input', updatePreview);