// Function to populate the product dropdown
export function getProductDropdown() {
    fetch(`http://localhost:8080/api/v1/product/list`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            const products = data.data;
            const idNameSelect = document.getElementById('idName');
            idNameSelect.innerHTML = ''; // Clear the dropdown before appending

            products.forEach(product => {
                const option = document.createElement("option");
                option.value = product.id;  // Use product ID as the value
                option.text = product.name; // Use product name as the text
                idNameSelect.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Error fetching product dropdown:', error);
        });
}

document.addEventListener("DOMContentLoaded", function () {
    getProductDropdown(); // Populate the product dropdown when the page loads
});