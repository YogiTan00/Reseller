document.getElementById('create-transaction').addEventListener('click', createTransaction);

export function createTransaction() {
        const idName = document.getElementById('idName').value;
        const name = document.getElementById('name').value;
        const returnItem = document.getElementById('returnItem').value;
        const dateNow = document.getElementById('dateNow').value;
        const unit = document.getElementById('unit').value;
        const description = document.getElementById('description').value;

        fetch('http://localhost:8080/api/v1/transaction/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                idName: idName,
                name: name,
                returnItem: returnItem,
                salesDate: dateNow,
                unit: unit,
                description: description,
            })
        })
            .then(response => response.json())
            .then(data => alert('Transaction created with Name: ' + data.name)) // Adjust based on the actual response structure
            .catch(error => console.error('Error:', error));
}