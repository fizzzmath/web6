rows = document.querySelectorAll('tbody > tr');

rows.forEach((row) => {
    row.addEventListener('click', () => {
        id = row.querySelector('td:first-child').innerText;

        window.location = `main.cgi?action=form&id=${id}`;
    });
});