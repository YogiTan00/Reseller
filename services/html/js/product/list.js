export function getProductList() {
    fetch('http://localhost:8080/api/v1/product/list', { // Adjust this endpoint based on your API
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            const products = data.data;
            const productList = document.getElementById('product-list');

            // Clear the product list container before adding new items
            productList.innerHTML = '';

            // Iterate over each product and create a list item
            products.forEach(product => {
                const productImage = product.image && product.image !== '' ? product.image : "";

                const productItem = `
                        <li class="list-group-item d-flex align-items-center">
                            <img src="${productImage}" alt="${product.name}">
                            <div>
                                <h5>${product.name}</h5>
                                <p>Size: ${product.typeSize}</p>
                                <p>Price: Rp ${product.price}</p>
                                <p>Default Price: Rp ${product.defaultPrice}</p>
                            </div>
                        </li>
                    `;
                // Insert the list item into the product list
                productList.insertAdjacentHTML('beforeend', productItem);
            });
        })
        .catch(error => {
            console.error('Error fetching products:', error);
        });
}