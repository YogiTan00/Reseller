import { createProduct } from '/product/create.js';
import { getProductList } from '/product/list.js';

document.addEventListener('DOMContentLoaded', function() {
    createProduct(); // Initialize product creation functionality
    getProductList(); // Initialize product list retrieval
});