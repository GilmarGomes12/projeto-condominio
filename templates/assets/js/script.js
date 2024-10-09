// Aplicar máscaras
$(document).ready(function() {
    $('.data-nascimento').inputmask('99/99/9999');
    $('.telefone').inputmask('(99)99999-9999');
    $('.rg').inputmask('99.999.999-9');
    $('.cpf').inputmask('999.999.999-99');
    $('.cep').inputmask('99999-999');
    $('.placa').inputmask('AAA-9999');
});

// Adiciona um novo campo de morador
document.getElementById('add-morador').addEventListener('click', function() {
    const moradoresDiv = document.getElementById('moradores');
    const newMoradorDiv = document.createElement('div');
    newMoradorDiv.classList.add('morador', 'mb-3');
    newMoradorDiv.innerHTML = `
        <input type="text" class="form-control" name="nome_morador[]" placeholder="Nome Completo" required>
        <input type="date" class="form-control mt-2 data-nascimento" name="data_nascimento[]" placeholder="Data de Nascimen
        to" required>
    `;
    moradoresDiv.appendChild(newMoradorDiv);
});

// Aplique a máscara de entrada a qualquer campo com a classe 'data-nascimento', mesmo que seja adicionado dinamicamente
$(document).on('input', '.data-nascimento', function() {
    $(this).inputmask('99/99/9999');
});

// Adiciona um novo campo de veículo
document.getElementById('add-veiculo').addEventListener('click', function() {
    const veiculosDiv = document.getElementById('veiculos');
    const newVeiculoDiv = document.createElement('div');
    newVeiculoDiv.classList.add('veiculo', 'mb-3');
    newVeiculoDiv.innerHTML = `
        <input type="text" class="form-control" name="placa[]" placeholder="Placa">
        <input type="text" class="form-control mt-2" name="cor[]" placeholder="Cor">
        <input type="text" class="form-control mt-2" name="marca[]" placeholder="Marca">
        <input type="text" class="form-control mt-2" name="modelo[]" placeholder="Modelo">
    `;
    veiculosDiv.appendChild(newVeiculoDiv);
});

// Validação de email, RG e CPF
document.getElementById('register-form').addEventListener('submit', function(event) {
    const emails = document.querySelectorAll('input[type="email"]');
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    for (let email of emails) {
        if (!emailRegex.test(email.value)) {
            alert('Por favor, insira um email válido.');
            event.preventDefault();
            return;
        }
    }

    const rgs = document.querySelectorAll('.rg');
    const rgRegex = /^\d{2}\.\d{3}\.\d{3}-\d{1}$/;
    for (let rg of rgs) {
        if (!rgRegex.test(rg.value)) {
            alert('Por favor, insira um RG válido.');
            event.preventDefault();
            return;
        }
    }

    const cpfs = document.querySelectorAll('.cpf');
    const cpfRegex = /^\d{3}\.\d{3}\.\d{3}-\d{2}$/;
    for (let cpf of cpfs) {
        if (!cpfRegex.test(cpf.value)) {
            alert('Por favor, insira um CPF válido.');
            event.preventDefault();
            return;
        }
    }
});