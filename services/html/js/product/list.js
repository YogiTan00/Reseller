let currentPage = 1; // Keep track of the current page
const limit = 10; // Set the number of products per page
let orderBy = 'created_at'
let sort = 'asc';

export function getProductList(page = 1, q = '', orderBy = '', sort = '') {
    const offset = (page - 1) * limit; // Calculate offset based on page
    const queryParams = new URLSearchParams({
        q: q,
        orderBy: orderBy, // Assuming you have a variable 'orderBy'
        limit: limit,
        offset: offset, // Use the calculated offset
        sort: sort
    });
    fetch(`http://localhost:8080/api/v1/product/list?${queryParams}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            const products = data.data;
            const productList = document.getElementById('product-list');
            const currentPageDisplay = document.getElementById('current-page');

            // Clear the product list container
            productList.innerHTML = '';

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
            products.forEach(product => {
                const productImage = product.image && product.image !== '' ? product.image : "https://via.placeholder.com/150";

                const productItem = `
                <li class="col-12 list-group-item d-flex align-items-center">
                    <div class="col-sm-1">
                        <img src="${productImage}">
                    </div>
                    <div class="col-sm-6">
                        <h5>${product.name}</h5> 
                    </div>
                    <div class="col-sm-1">
                        <p>Size: ${product.typeSize}</p>
                    </div>
                    <div class="col-sm-2 col-price">
                        <p>Price: Rp ${product.price}</p>
                        <p>Default Price: Rp ${product.defaultPrice}</p>
                    </div>
                    <div class="col-sm-2 col-update">
                        <p>Edit</p>
                        <p>Delete</p>
                    </div>
                </li>
            `;
                productList.insertAdjacentHTML('beforeend', productItem);
            });
        })
        .catch(error => {
            console.error('Error fetching products:', error);
        });
}

// Add event listener to the "Search" button
document.getElementById('apply-filter').addEventListener('click', () => {
    const search = document.getElementById('q').value; // Use 'q' for input ID
    currentPage = 1;  // Reset to page 1 when searching
    getProductList(currentPage, search, orderBy ,sort);
});

// Add event listener to the "Next" button
document.getElementById('next-page').addEventListener('click', () => {
    currentPage += 1;
    const search = document.getElementById('q').value; // Use 'q' for input ID
    getProductList(currentPage, search, orderBy ,sort);
});

// Add event listener to the "Previous" button
document.getElementById('prev-page').addEventListener('click', () => {
    if (currentPage > 1) {
        currentPage -= 1;
        const search = document.getElementById('q').value; // Use 'q' for input ID
        getProductList(currentPage, search, orderBy ,sort);
    }
});

document.getElementById('sort').addEventListener('click', function() {
    sort = (sort === 'asc') ? 'desc' : 'asc';
    const search = document.getElementById('q').value; // Use 'q' for input ID
    getProductList(currentPage, search, orderBy ,sort);
});

// Initially fetch the product list for the first page
getProductList(currentPage);
