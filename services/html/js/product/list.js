export function getProductList() {
    fetch('/api/v1/product/list', { // Adjust this endpoint based on your API
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            // Assuming `data` is an array of products
            const productListContainer = document.getElementById('productList');
            productListContainer.innerHTML = ''; // Clear previous list
            data.forEach(product => {
                const productItem = document.createElement('li');
                productItem.textContent = `Name: ${product.name}, Type Size: ${product.typeSize}, Price: ${product.price}`;
                productListContainer.appendChild(productItem);
            });
        })
        .catch(error => console.error('Error fetching product list:', error));
}