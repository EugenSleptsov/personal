   (function() {
        function sortTable(n, table) {
            const tbody = table.tBodies[0];
            const rows = Array.from(tbody.querySelectorAll('tr'));

            const compare = (a, b) => {
                const aValue = parseFloat(a.dataset.money);
                const bValue = parseFloat(b.dataset.money);

                if (aValue === bValue) return 0;
                return aValue < bValue ? -1 : 1;
            };

            const sortedRows = rows.sort((a, b) => {
                const tdA = a.querySelectorAll('td')[n];
                const tdB = b.querySelectorAll('td')[n];
                return compare(tdA, tdB);
            });

            const th = table.querySelector(`th:nth-child(${n + 1})`);
            const otherHeaders = Array.from(table.querySelectorAll(`th:not(:nth-child(${n + 1}))`));
            otherHeaders.forEach(header => {
                header.classList.remove('sorted-asc');
                header.classList.remove('sorted-desc');
            });

            if (th.classList.contains('sorted-desc')) {
                th.classList.remove('sorted-desc');
                th.classList.add('sorted-asc');
            } else {
                sortedRows.reverse();
                th.classList.add('sorted-desc');
                th.classList.remove('sorted-asc');
            }

            tbody.innerHTML = '';
            sortedRows.forEach(row => tbody.appendChild(row));
        }

        document.addEventListener('DOMContentLoaded', () => {
            const table = document.getElementById('debt_table');
            const allHeaders = Array.from(table.querySelectorAll('th'));

            table.addEventListener('click', (event) => {
                const header = event.target.closest('th.sortable');
                if (!header) return;
                const index = allHeaders.indexOf(header);
                sortTable(index, table);
            });
        });
    })();
