// Script to handle image file name display
const imageInput = document.getElementById('image');
if (imageInput) {
    imageInput.addEventListener('change', function (event) {
        const fileInput = event.target;
        const fileName = fileInput.files[0] ? fileInput.files[0].name : 'No file chosen';

        // Update the displayed file name
        console.log(fileName)
        document.getElementById('file-name').textContent = fileName;
    });
}

// Get the current date in yyyy-mm-dd format
const today = new Date().toISOString().split('T')[0];
const dateInput = document.getElementById('dateNow');
if (dateInput !== null) {
    dateInput.value = today;
}

// Get the input and buttons
const unitInput = document.getElementById('unit');
const increaseBtn = document.getElementById('increaseUnit');
const decreaseBtn = document.getElementById('decreaseUnit');

// // Increase unit value
// increaseBtn.addEventListener('click', function() {
//     if (unitInput.value > 0) {
//         unitInput.value = parseInt(unitInput.value) + 1;
//     }
// });
//
// // Decrease unit value, ensuring it doesn't go below 0
// decreaseBtn.addEventListener('click', function() {
//     if (unitInput.value > 0) {
//         unitInput.value = parseInt(unitInput.value) - 1;
//     }
// });

function DateRangePickerById() {
    const input = document.getElementById('date-range');
    if (input) {
        $(input).daterangepicker({
            opens: 'left',
            locale: {
                format: 'YYYY-MM-DD'
            }
        });

        $(input).on('apply.daterangepicker', function(ev, picker) {
            input.value = picker.startDate.format('YYYY-MM-DD') + ' - ' + picker.endDate.format('YYYY-MM-DD');
        });


        $(input).on('cancel.daterangepicker', function(ev, picker) {
            input.value = '';
        });
    }
}

export async function loadHTML(id, url) {
    const container = document.getElementById(id);
    if (!container) {
        console.warn(`Element with id "${id}" not found.`);
        return;
    }

    try {
        const res = await fetch(url);
        if (!res.ok) throw new Error(`Failed to load ${url}: ${res.statusText}`);
        const html = await res.text();
        container.innerHTML = html;
        
        // Load dropdown after HTML is loaded
        if (id === 'create-single-transaction') {
            console.log('Single transaction HTML loaded, calling loadProductDropdown...');
            setTimeout(() => {
                if (window.loadProductDropdown) {
                    window.loadProductDropdown();
                }
            }, 100);
        } else if (id === 'create-multiple-transaction') {
            console.log('Multiple transaction HTML loaded, calling loadProductDropdown...');
            setTimeout(() => {
                if (window.loadProductDropdown) {
                    window.loadProductDropdown();
                }
                if (window.populateDropdowns) {
                    window.populateDropdowns();
                }
            }, 100);
        }
    } catch (error) {
        container.innerHTML = `<p class="text-danger">${error.message}</p>`;
    }
}

function loadAll() {
    // Always load HTML content
    console.log('loadAll called - loading transaction forms...');
    
    console.log('Loading single transaction HTML...');
    loadHTML('create-single-transaction', '/html/transaction/single.html');
    
    console.log('Loading multiple transaction HTML...');
    loadHTML('create-multiple-transaction', '/html/transaction/multiple.html');
}


window.addEventListener('DOMContentLoaded', loadAll);
document.addEventListener("DOMContentLoaded", function () {
    DateRangePickerById();
});

// Load dropdown after dynamic content is loaded
window.addEventListener('load', function() {
    console.log('Window fully loaded, calling populateDropdowns...');
    if (window.populateDropdowns) {
        window.populateDropdowns();
    }
});
