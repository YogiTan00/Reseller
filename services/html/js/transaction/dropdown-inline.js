console.log('Dropdown inline script loaded');

window.allProducts = [];

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
        populateSingleDropdown();
        populateMultipleDropdowns();
    })
    .catch(error => {
        console.error('Error fetching product dropdown:', error);
        alert('Gagal memuat daftar produk. Silakan refresh halaman.');
    });

function populateSingleDropdown() {
    const idNameSelect = document.getElementById('idName');
    if (!idNameSelect) return;

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
        option.text = product.name + ' (' + product.type_size + ')';
        idNameSelect.appendChild(option);
    });

    idNameSelect.addEventListener("change", function() {
        const selectedId = idNameSelect.value;
        const selectedProduct = window.allProducts.find(product => product.id === selectedId);
        if (selectedProduct) {
            const productNameInput = document.getElementById('productName');
            if (productNameInput) {
                productNameInput.value = selectedProduct.name;
            }
        }
    });
}

function populateMultipleDropdowns() {
    const dropdowns = document.querySelectorAll('.id-name-multiple');
    dropdowns.forEach(dropdown => {
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
            option.text = product.name + ' (' + product.type_size + ')';
            dropdown.appendChild(option);
        });
    });
}

window.populateDropdowns = function() {
    populateMultipleDropdowns();
};
