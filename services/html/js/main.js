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