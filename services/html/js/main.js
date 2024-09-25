import { createProduct } from 'product/create.js';
import { getProductList } from 'product/list.js';

document.addEventListener('DOMContentLoaded', function() {
    // Call createProduct when the relevant event occurs (e.g., form submission)
    const createButton = document.getElementById('createProduct'); // Replace with your actual button ID
    if (createButton) {
        createButton.addEventListener('click', createProduct);
    }

    // Call getProductList to initialize product list retrieval
    getProductList();
});

// Script to handle image file name display
document.getElementById('image').addEventListener('change', function(event) {
    const fileInput = event.target;
    const fileName = fileInput.files[0] ? fileInput.files[0].name : 'No file chosen';

    // Update the displayed file name
    console.log(fileName)
    document.getElementById('file-name').textContent = fileName;
});