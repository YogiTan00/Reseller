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

            // Create and append a default option
            const defaultOption = document.createElement("option");
            defaultOption.value = ''; // Set an empty value for the default option
            defaultOption.text = 'Select a Product'; // Default text
            defaultOption.disabled = true; // Make it unselectable
            defaultOption.selected = true; // Make it selected by default
            idNameSelect.appendChild(defaultOption);

            products.forEach(product => {
                const option = document.createElement("option");
                option.value = product.id;  // Use product ID as the value
                option.text = product.name; // Use product name as the text
                idNameSelect.appendChild(option);
            });
            idNameSelect.addEventListener("change", () => {
                const selectedId = idNameSelect.value;
                const selectedProduct = products.find(product => product.id === selectedId);
                if (selectedProduct) {
                    document.getElementById('productName').value = selectedProduct.name;
                } else {
                    document.getElementById('productName').value = '';
                }
            });
        })
        .catch(error => {
            console.error('Error fetching product dropdown:', error);
        });
}

document.addEventListener("DOMContentLoaded", function () {
    getProductDropdown(); // Populate the product dropdown when the page loads
});