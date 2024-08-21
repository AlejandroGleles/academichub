document.addEventListener('DOMContentLoaded', function() {
    aDashboard();
    listarA();
    cAluno(); // Inicializa o cadastro de alunos
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



function cAluno() {
    const alunoForm = document.getElementById('alunoForm');
    if (alunoForm) {
        alunoForm.addEventListener('submit', function(e) {
            e.preventDefault();

            const nome = document.getElementById('nome').value;
            const matricula = document.getElementById('matricula').value;
            const turma = document.getElementById('turma').value;

            // Valida se a matrícula é um número inteiro
            if (!Number.isInteger(parseInt(matricula))) {
                alert('A matrícula deve ser um número inteiro.');
                return;
            }

            const aluno = {
                Nome: nome,
                Matricula: parseInt(matricula), // Converte para inteiro
                Turma: turma,
            };

            fetch('http://localhost:8080/api/v1/aluno', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(aluno)
            })
            .then(response => response.json())
            .then(data => {
                console.log('Success:', data);
                alert('Aluno cadastrado com sucesso!');
                window.location.href = 'principal.html'; // Ajuste o nome da página inicial conforme necessário
            })
            .catch((error) => {
                console.error('Error:', error);
                alert('Erro ao cadastrar aluno');
            });
        });
    } else {
        console.error('Formulário de cadastro de alunos não encontrado');
    }
}





function listarA() {
    fetch('http://localhost:8080/api/v1/alunos', {
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
        const listaDiv = document.getElementById('alunoLista');
        listaDiv.innerHTML = '';

        if (data.data && Array.isArray(data.data)) {
            data.data.forEach(aluno => {
                const alunoCard = document.createElement('div');
                alunoCard.className = 'col-md-4';
                alunoCard.dataset.id = aluno.ID;
                alunoCard.innerHTML = `
                    <div class="card mb-3">
                        <div class="card-body">
                            <h5 class="card-title">${aluno.Nome}</h5>
                            <p class="card-text"><strong>Matrícula:</strong> ${aluno.Matricula}</p>
                            <p class="card-text"><strong>Turma:</strong> ${aluno.Turma}</p>
                            <div class="card-actions">
                                <button class="btn btn-warning" onclick="editarAluno(${aluno.ID}, '${encodeURIComponent(aluno.Nome)}', '${encodeURIComponent(aluno.Matricula)}', '${encodeURIComponent(aluno.Turma)}')">Editar</button>
                                <button class="btn btn-danger" onclick="deletarAluno(${aluno.ID})">Deletar</button>
                            </div>
                        </div>
                    </div>
                `;
                listaDiv.appendChild(alunoCard);
            });
        } else {
            listaDiv.innerHTML = '<p>No data available.</p>';
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        document.getElementById('alunoLista').innerHTML = '<p>Failed to load data.</p>';
    });
}

function editarAluno(id, nome, matricula, turma) {
    const url = `edita_aluno.html?id=${id}&nome=${nome}&matricula=${matricula}&turma=${turma}`;
    window.location.href = url;
}

function deletarAluno(id) {
    fetch(`http://localhost:8080/api/v1/aluno?id=${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Erro ao deletar aluno');
        }
        return response.json();
    })
    .then(data => {
        console.log('Success:', data);
        alert('Aluno deletado com sucesso!');
        listarA(); // Atualiza a lista de alunos
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Erro ao deletar aluno');
    });
}
