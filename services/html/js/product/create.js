document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('productForm').addEventListener('submit', function(event) {
        event.preventDefault();

        const name = document.getElementById('name').value;
        const typeSize = document.getElementById('typeSize').value;
        const marketType = document.getElementById('marketType').value;
        const image = document.getElementById('image').value;
        const defaultPrice = document.getElementById('defaultPrice').value;
        const price = document.getElementById('price').value;

        fetch('/api/v1/product/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: name,
                typeSize: typeSize,
                marketType: marketType,
                image: image,
                defaultPrice: defaultPrice,
                price: price
            })
        })
            .then(response => response.json())
            .then(data => alert('Product created with Name: ' + data.name)) // Adjust based on the actual response structure
            .catch(error => console.error('Error:', error));
    });
});