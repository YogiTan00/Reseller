// Script to handle image file name display
const imageInput = document.getElementById('image');
if (imageInput) {
    imageInput.addEventListener('change', function(event) {
        const fileInput = event.target;
        const fileName = fileInput.files[0] ? fileInput.files[0].name : 'No file chosen';

        // Update the displayed file name
        console.log(fileName)
        document.getElementById('file-name').textContent = fileName;
    });
}

// Get the current date in yyyy-mm-dd format
const today = new Date().toISOString().split('T')[0];
    document.getElementById('dateNow').value = today;

// Get the input and buttons
const unitInput = document.getElementById('unit');
const increaseBtn = document.getElementById('increaseUnit');
const decreaseBtn = document.getElementById('decreaseUnit');

// Increase unit value
increaseBtn.addEventListener('click', function() {
    unitInput.value = parseInt(unitInput.value) + 1;
});

// Decrease unit value, ensuring it doesn't go below 0
decreaseBtn.addEventListener('click', function() {
    if (unitInput.value > 0) {
        unitInput.value = parseInt(unitInput.value) - 1;
    }
});