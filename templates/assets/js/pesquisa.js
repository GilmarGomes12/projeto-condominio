const form = document.getElementById('pesquisar-form');
const resultadosContainer = document.querySelector('.resultados-container');
const listaResultados = document.getElementById('lista-resultados');
const semResultados = document.getElementById('sem-resultados');

form.addEventListener('submit', (event) => {
    event.preventDefault();

    const query = document.getElementById('query').value;
    const tipoPesquisa = document.getElementById('tipo-pesquisa').value;

    fetch(`/pesquisar?query=${query}&tipo=${tipoPesquisa}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erro na requisição: ${response.statusText}`);
            }
            return response.json();
        })
        .then(data => {
            console.log('Dados recebidos:', data); // Log dos dados recebidos para depuração
            listaResultados.innerHTML = ''; // Limpa a lista de resultados
            if (data.length > 0) {
                resultadosContainer.style.display = 'block';
                semResultados.style.display = 'none';

                data.forEach(resultado => {
                    const li = document.createElement('li');
                    li.classList.add('list-group-item');
                    li.textContent = resultado.nome || resultado.descricao || resultado.titulo;
                    listaResultados.appendChild(li);
                });
            } else {
                resultadosContainer.style.display = 'block';
                semResultados.style.display = 'block';
            }
        })
        .catch(error => {
            console.error('Erro ao buscar resultados:', error);
            semResultados.textContent = 'Erro ao buscar resultados. Por favor, tente novamente.';
            resultadosContainer.style.display = 'block';
            semResultados.style.display = 'block';
        });
});