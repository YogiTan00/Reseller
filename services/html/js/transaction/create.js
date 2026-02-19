document.getElementById('create-transaction').addEventListener('click', createTransaction);

export async function createTransaction() {
    const idName = document.getElementById('idName').value;
    const returnItem = document.getElementById('returnItem').value;
    const dateNow = document.getElementById('dateNow').value;
    const unit = document.getElementById('unit').value;
    const description = document.getElementById('description').value.trim();

    if (!idName) {
        alert('Mohon pilih produk!');
        return;
    }

    if (!dateNow) {
        alert('Mohon pilih tanggal!');
        return;
    }

    if (!unit || parseInt(unit) <= 0) {
        alert('Jumlah harus lebih dari 0!');
        return;
    }

    const btn = document.getElementById('create-transaction');
    const originalText = btn.innerHTML;
    btn.innerHTML = '<i class="bi bi-spinner"></i> Menyimpan...';
    btn.disabled = true;

    const body = {
        idName: idName,
        returnItem: returnItem === 'true',
        salesDate: dateNow,
        unit: parseInt(unit),
        description: description
    };

    try {
        const res = await axios.post('http://localhost:8080/api/v1/transaction/create', body);
        alert('Transaksi berhasil ditambahkan!');
        document.getElementById('transaction-form').reset();
        btn.innerHTML = originalText;
        btn.disabled = false;
        getTransactionList();
    } catch (err) {
        console.error(err);
        const errorMessage = err?.response?.data?.message || 'Gagal menambahkan transaksi. Silakan coba lagi.';
        alert(errorMessage);
        btn.innerHTML = originalText;
        btn.disabled = false;
    }
}

document.getElementById('create-multiple-transaction').addEventListener('click', createMultipleTransaction);

export async function createMultipleTransaction() {
    const items = document.querySelectorAll('.transaction-item');
    const transactions = [];

    for (let item of items) {
        const idName = item.querySelector('.id-name-multiple').value;
        const returnItem = item.querySelector('.return-item').value;
        const dateNow = item.querySelector('.date-now').value;
        const unit = item.querySelector('.unit').value;
        const description = item.querySelector('.description').value.trim();

        if (!idName) {
            alert('Mohon pilih produk untuk semua item!');
            return;
        }

        if (!dateNow) {
            alert('Mohon pilih tanggal untuk semua item!');
            return;
        }

        if (!unit || parseInt(unit) <= 0) {
            alert('Jumlah harus lebih dari 0 untuk semua item!');
            return;
        }

        transactions.push({
            idName: idName,
            returnItem: returnItem === 'true',
            salesDate: dateNow,
            unit: parseInt(unit),
            description: description
        });
    }

    if (transactions.length === 0) {
        alert('Tidak ada transaksi untuk disimpan!');
        return;
    }

    const btn = document.getElementById('create-multiple-transaction');
    const originalText = btn.innerHTML;
    btn.innerHTML = '<i class="bi bi-spinner"></i> Menyimpan...';
    btn.disabled = true;

    try {
        for (let transaction of transactions) {
            await axios.post('http://localhost:8080/api/v1/transaction/create', transaction);
        }
        alert('Semua transaksi berhasil ditambahkan!');
        document.getElementById('transaction-multiple-form').reset();
        document.getElementById('transaction-items').innerHTML = '';
        addInitialTransactionItem();
        btn.innerHTML = originalText;
        btn.disabled = false;
        getTransactionList();
    } catch (err) {
        console.error(err);
        const errorMessage = err?.response?.data?.message || 'Gagal menambahkan transaksi. Silakan coba lagi.';
        alert(errorMessage);
        btn.innerHTML = originalText;
        btn.disabled = false;
    }
}

let itemCount = 0;

window.addTransactionItem = function() {
    itemCount++;
    const newItem = document.createElement('div');
    newItem.className = 'transaction-item';
    newItem.id = `transaction-item-${itemCount}`;
    newItem.innerHTML = `
        <div class="transaction-item-header">
            <span class="transaction-item-number">Item #${itemCount}</span>
            <div class="action-buttons">
                <button type="button" class="btn-add" onclick="addTransactionItem()">
                    <i class="bi bi-plus"></i> Tambah
                </button>
                <button type="button" class="btn-remove" onclick="removeTransactionItem(${itemCount})">
                    <i class="bi bi-trash"></i> Hapus
                </button>
            </div>
        </div>
        <div class="row g-3">
            <div class="col-md-5">
                <label class="form-label">Pilih Produk</label>
                <select class="form-select id-name-multiple" required>
                    <option value="" disabled selected>-- Pilih Produk --</option>
                </select>
                <input type="hidden" class="product-name-multiple">
            </div>
            <div class="col-md-2">
                <label class="form-label">Status Retur</label>
                <select class="form-select return-item">
                    <option value="false" selected>Penjualan</option>
                    <option value="true">Retur</option>
                </select>
            </div>
            <div class="col-md-2">
                <label class="form-label">Tanggal</label>
                <input type="date" class="form-control date-now" required>
            </div>
            <div class="col-md-1">
                <label class="form-label">Jumlah</label>
                <input type="number" class="form-control unit" value="1" min="1" required>
            </div>
            <div class="col-md-2">
                <label class="form-label">Keterangan</label>
                <input type="text" class="form-control description" placeholder="...">
            </div>
        </div>
    `;
    document.getElementById('transaction-items').appendChild(newItem);
    
    setTimeout(() => {
        const allDropdowns = document.querySelectorAll('.id-name-multiple');
        if (typeof window.populateDropdowns === 'function') {
            window.populateDropdowns();
        } else if (window.allProducts && window.allProducts.length > 0) {
            allDropdowns.forEach(dropdown => {
                dropdown.innerHTML = '<option value="" disabled selected>-- Pilih Produk --</option>';
                window.allProducts.forEach(product => {
                    const option = document.createElement("option");
                    option.value = product.id;
                    const name = product.name || 'Unknown Product';
                    const typeSize = product.typeSize || '-';
                    option.text = `${name} (${typeSize})`;
                    dropdown.appendChild(option);
                });
            });
        }
    }, 100);
};

window.removeTransactionItem = function(id) {
    const item = document.getElementById(`transaction-item-${id}`);
    if (item) {
        item.remove();
    }
};

function addInitialTransactionItem() {
    if (document.getElementById('transaction-items').children.length === 0) {
        addTransactionItem();
    }
}
