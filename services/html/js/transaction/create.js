document.getElementById('create-transaction').addEventListener('click', createTransaction);

export async function createTransaction() {
    const idName = document.getElementById('idName').value;
    const name = document.getElementById('productName').value;
    const returnItem = document.getElementById('returnItem').value;
    const dateNow = document.getElementById('dateNow').value;
    const unit = document.getElementById('unit').value;
    const description = document.getElementById('description').value;
    const body = {
            idName: idName,
            name: name,
            ...(returnItem ? {returnItem: JSON.parse(returnItem)} : {}),
            salesDate: dateNow,
            unit: unit,
            description: description,
        }
    axios.post('http://localhost:8080/api/v1/transaction/create',body)
        .then((res)=>{
            console.log(res)
        })
        .catch((err)=> alert(err?.response?.data?.message))
}