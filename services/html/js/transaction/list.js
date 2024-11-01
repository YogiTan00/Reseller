let currentPage = 1; // Keep track of the current page
const limit = 10; // Set the number of products per page
let orderBy = 'created_at'
let sort = 'asc';

export function getTransactionList(page = 1, q = '', orderBy = '', sort = '') {
    const offset = (page - 1) * limit; // Calculate offset based on page
    const queryParams = new URLSearchParams({
        q: q,
        orderBy: orderBy, // Assuming you have a variable 'orderBy'
        limit: limit,
        offset: offset, // Use the calculated offset
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

            // Clear the product list container
            transactionList.innerHTML = '';

            // Update the current page display
            currentPageDisplay.textContent = currentPage;

            // Check the total count of products and set pagination button states
            const totalCount = data.meta.count; // Total number of products
            const totalPages = Math.ceil(totalCount / limit); // Calculate total pages

            // Enable or disable pagination buttons based on the current page
            const nextPageButton = document.getElementById('next-page');
            const prevPageButton = document.getElementById('prev-page');

            nextPageButton.disabled = currentPage >= totalPages; // Disable next if on last page
            prevPageButton.disabled = currentPage <= 1; // Disable previous if on first page

            // Render product items
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
            console.error('Error fetching products:', error);
        });
}

// Add event listener to the "Search" button
const applyFilterBtn = document.getElementById('apply-filter');
if (applyFilterBtn) {
    applyFilterBtn.addEventListener('click', () => {
        const search = document.getElementById('q').value; // Use 'q' for input ID
        currentPage = 1;  // Reset to page 1 when searching
        getTransactionList(currentPage, search, orderBy ,sort);
    });
}

// Add event listener to the "Next" button
const nextPagerBtn = document.getElementById('next-page')
if (nextPagerBtn) {
    nextPagerBtn.addEventListener('click', () => {
        currentPage += 1;
        const search = document.getElementById('q').value; // Use 'q' for input ID
        getTransactionList(currentPage, search, orderBy, sort);
    });
}

// Add event listener to the "Previous" button
const prevPagerBtn = document.getElementById('prev-page')
if (prevPagerBtn) {
    prevPagerBtn.addEventListener('click', () => {
        if (currentPage > 1) {
            currentPage -= 1;
            const search = document.getElementById('q').value; // Use 'q' for input ID
            getTransactionList(currentPage, search, orderBy ,sort);
        }
    });
}

const sortBtn = document.getElementById('sort');
if (sortBtn) {
    sortBtn.addEventListener('click', function() {
        sort = (sort === 'asc') ? 'desc' : 'asc';
        const search = document.getElementById('q').value; // Use 'q' for input ID
        getTransactionList(currentPage, search, orderBy ,sort);
    });
}

// Call this function after the DOM is fully loaded
document.addEventListener('DOMContentLoaded', () => {
    getTransactionList(); // Call the function on load to fetch the product list
});