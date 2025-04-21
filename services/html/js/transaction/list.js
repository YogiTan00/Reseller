let currentPage = 1; // Keep track of the current page
const limit = 10; // Set the number of transactions per page
let orderBy = 'created_at';
let sort = 'asc';

export function getTransactionList(page = 1, q = '') {
    const offset = (page - 1) * limit; // Calculate offset based on page
    const queryParams = new URLSearchParams({
        q: q,
        orderBy: orderBy,
        limit: limit,
        offset: offset,
        sort: sort
    });
    fetch(`http://localhost:8080/api/v1/transaction/list?${queryParams}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            const transactions = data.data;
            const transactionList = document.getElementById('transaction-list');
            const currentPageDisplay = document.getElementById('current-page');

            // Clear the transaction list container
            transactionList.innerHTML = '';

            // Update the current page display
            currentPageDisplay.textContent = currentPage;

            // Check the total count of transactions and set pagination button states
            const totalCount = data.meta ? data.meta.count : transactions.length;
            const totalPages = Math.ceil(totalCount / limit); // Calculate total pages

            // Enable or disable pagination buttons based on the current page
            const nextPageButton = document.getElementById('next-page');
            const prevPageButton = document.getElementById('prev-page');

            nextPageButton.disabled = currentPage >= totalPages;
            prevPageButton.disabled = currentPage <= 1;

            // Render transaction items
            transactions.forEach(transaction => {
                const date = new Date(transaction.salesDate);
                const transactionItem = `
                <li class="col-12 list-group-item d-flex align-items-center">
                    <div class="col-sm-4">
                        <h5>${transaction.name}</h5> 
                    </div>
                    <div class="col-sm-1">
                        <p>Return: ${transaction.returnItem}</p>
                    </div>
                    <div class="col-sm-3">
                        <p>Sales Date: ${date.toLocaleDateString("id-CA")}</p>
                    </div>
                    </div>
                    <div class="col-sm-1">
                        <p>Unit: ${transaction.unit}</p>
                    </div>
                    <div class="col-sm-2">
                        <p>Description: ${transaction.description}</p>
                    </div>
                    <div class="col-sm-1 col-update">
                        <p>Edit</p>
                        <p>Delete</p>
                    </div>
                </li>
            `;
                transactionList.insertAdjacentHTML('beforeend', transactionItem);
            });
        })
        .catch(error => {
            console.error('Error fetching transactions:', error);
        });
}

// Add event listeners for pagination, filtering, and sorting
const applyFilterBtn = document.getElementById('apply-filter');
if (applyFilterBtn) {
    applyFilterBtn.addEventListener('click', () => {
        const search = document.getElementById('q').value;
        currentPage = 1;  // Reset to page 1 when searching
        getTransactionList(currentPage, search);
    });
}

const nextPagerBtn = document.getElementById('next-page');
if (nextPagerBtn) {
    nextPagerBtn.addEventListener('click', () => {
        currentPage += 1;
        const search = document.getElementById('q').value;
        getTransactionList(currentPage, search);
    });
}

const prevPagerBtn = document.getElementById('prev-page');
if (prevPagerBtn) {
    prevPagerBtn.addEventListener('click', () => {
        if (currentPage > 1) {
            currentPage -= 1;
            const search = document.getElementById('q').value;
            getTransactionList(currentPage, search);
        }
    });
}

const sortBtn = document.getElementById('sort');
if (sortBtn) {
    sortBtn.addEventListener('click', () => {
        sort = (sort === 'asc') ? 'desc' : 'asc';
        const search = document.getElementById('q').value;
        getTransactionList(currentPage, search);
    });
}

// Call this function after the DOM is fully loaded
document.addEventListener('DOMContentLoaded', () => {
    getTransactionList(); // Fetch the transaction list on page load
});
