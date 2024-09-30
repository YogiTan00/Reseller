import { getProductList } from './product/list.js';
document.addEventListener('DOMContentLoaded', function() {
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