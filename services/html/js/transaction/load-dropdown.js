console.log('load-dropdown.js loaded');

fetch('http://localhost:8080/api/v1/product/list', {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
})
    .then(response => {
        console.log('Product list response status:', response.status);
        return response.json();
    })
    .then(data => {
        console.log('Product list data:', data);
        console.log('Number of products:', data.data ? data.data.length : 0);
        window.allProducts = data.data || [];
        
        console.log('Starting to populate dropdowns...');
        populateSingleDropdown();
        populateMultipleDropdowns();
    })
    .catch(error => {
        console.error('Error fetching product dropdown:', error);
        alert('Gagal memuat daftar produk. Silakan refresh halaman.');
    });

function populateSingleDropdown() {
    console.log('populateSingleDropdown called');
    const idNameSelect = document.getElementById('idName');
    if (!idNameSelect) {
        console.log('idName element not found');
        return;
    }
    console.log('idName element found, options before:', idNameSelect.options.length);

    idNameSelect.innerHTML = '';
    const defaultOption = document.createElement("option");
    defaultOption.value = '';
    defaultOption.text = '-- Pilih Produk --';
    defaultOption.disabled = true;
    defaultOption.selected = true;
    idNameSelect.appendChild(defaultOption);

    window.allProducts.forEach(product => {
        const option = document.createElement("option");
        option.value = product.id;
        const name = product.name || 'Unknown Product';
        const typeSize = product.typeSize || '-';
        option.text = `${name} (${typeSize})`;
        idNameSelect.appendChild(option);
    });

    idNameSelect.addEventListener("change", function() {
        const selectedId = idNameSelect.value;
        const selectedProduct = window.allProducts.find(function(product) { return product.id === selectedId; });
        if (selectedProduct) {
            const productNameInput = document.getElementById('productName');
            if (productNameInput) {
                productNameInput.value = selectedProduct.name;
            }
        }
    });
}

function populateMultipleDropdowns() {
    console.log('populateMultipleDropdowns called');
    const dropdowns = document.querySelectorAll('.id-name-multiple');
    console.log('Found dropdowns:', dropdowns.length);
    dropdowns.forEach(function(dropdown) {
        dropdown.innerHTML = '';
        const defaultOption = document.createElement("option");
        defaultOption.value = '';
        defaultOption.text = '-- Pilih Produk --';
        defaultOption.disabled = true;
        defaultOption.selected = true;
        dropdown.appendChild(defaultOption);

        window.allProducts.forEach(product => {
            const option = document.createElement("option");
            option.value = product.id;
            option.text = `${product.name} (${product.typeSize})`;
            dropdown.appendChild(option);
        });
    });
}
