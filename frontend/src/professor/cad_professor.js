document.addEventListener('DOMContentLoaded', function() {
    listarP();
    aDashboard();
    cProfessor(); // Inicializa o cadastro de professores
    setupUpdateForm(); // Configura o formulário de atualização
});

function aDashboard() {
    fetch('http://localhost:8080/api/v1/dashboard', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('totalProfessores').textContent = data.totalProfessores || 0;
        document.getElementById('totalTurmas').textContent = data.totalTurmas || 0;
        document.getElementById('totalAlunos').textContent = data.totalAlunos || 0;
        document.getElementById('totalAtividades').textContent = data.totalAtividades || 0;
    })
    .catch(error => {
        console.error('Erro ao carregar dados da dashboard:', error);
    });
}

function cProfessor() {
    const professorForm = document.getElementById('professorForm');
    if (professorForm) {
        professorForm.addEventListener('submit', function(e) {
            e.preventDefault();

            const nome = document.getElementById('nome').value;
            const email = document.getElementById('email').value;
            const cpf = document.getElementById('cpf').value;

            const professor = {
                Nome: nome,
                Email: email,
                CPF: cpf,
            };

            fetch('http://localhost:8080/api/v1/professor', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(professor)
            })
            .then(response => response.json())
            .then(data => {
                console.log('Success:', data);
                alert('Professor cadastrado com sucesso!');
                //listarP(); // Atualiza a lista de professores
                window.location.href = 'principal.html'; // Ajuste o nome da página inicial conforme necessário
            })
            .catch((error) => {
                console.error('Error:', error);
                alert('Erro ao cadastrar professor');
            });
        });
    } else {
        console.error('Formulário de cadastro de professores não encontrado');
    }
}

function listarP() {
    fetch('http://localhost:8080/api/v1/professores', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        const listaDiv = document.getElementById('professoresLista');
        listaDiv.innerHTML = '';

        if (data.data && Array.isArray(data.data)) {
            data.data.forEach(professor => {
                const professorCard = document.createElement('div');
                professorCard.className = 'col-md-4';
                professorCard.dataset.id = professor.ID;
                professorCard.innerHTML = `
                    <div class="card mb-3">
                        <div class="card-body">
                            <h5 class="card-title">${professor.Nome}</h5>
                            <p class="card-text"><strong>Identificação:</strong> ${professor.ID}</p>
                            <p class="card-text"><strong>E-mail:</strong> ${professor.Email}</p>
                            <p class="card-text"><strong>CPF:</strong> ${professor.CPF}</p>
                            <div class="card-actions">
                                <button class="btn btn-warning" onclick="editarProfessor(${professor.ID}, '${encodeURIComponent(professor.Nome)}', '${encodeURIComponent(professor.Email)}', '${encodeURIComponent(professor.CPF)}')">Editar</button>
                                <button class="btn btn-danger" onclick="deletarProfessor(${professor.ID})">Deletar</button>
                            </div>
                        </div>
                    </div>
                `;
                listaDiv.appendChild(professorCard);
            });
        } else {
            listaDiv.innerHTML = '<p>No data available.</p>';
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        document.getElementById('professoresLista').innerHTML = '<p>Failed to load data.</p>';
    });
}


function editarProfessor(id, nome, email, cpf) {
    const url = `edita_professor.html?id=${id}&nome=${nome}&email=${email}&cpf=${cpf}`;
    window.location.href = url;
}

function visualizarProfessor(id) {
    fetch(`http://localhost:8080/api/v1/professor?id=${id}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Erro ao buscar dados do professor');
        }
        return response.json();
    })
    .then(responseData => {
        // Acessa os dados do professor no objeto 'data'
        const professor = responseData.data;
        console.log('Dados do professor:', professor);

        // Exibindo os dados no formulário de edição
        document.getElementById('updateProfessorId').value = professor.ID;
        document.getElementById('updateNome').value = professor.Nome;
        document.getElementById('updateEmail').value = professor.Email;
        document.getElementById('updateCPF').value = professor.CPF;

        // Exibindo o formulário de edição
        document.getElementById('updateFormContainer').style.display = 'block';
    })
    .catch(error => {
        console.error('Erro ao carregar dados do professor para edição:', error);
        alert('Erro ao carregar dados do professor para edição');
    });
}



function deletarProfessor(id) {
    fetch(`http://localhost:8080/api/v1/professor?id=${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Erro ao deletar professor');
        }
        return response.json();
    })
    .then(data => {
        console.log('Success:', data);
        alert('Professor deletado com sucesso!');
        listarP(); // Atualiza a lista de professores
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Erro ao deletar professor');
    });
}
