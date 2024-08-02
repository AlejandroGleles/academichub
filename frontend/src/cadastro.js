//cadastra Professor
function cProfessor(){
document.getElementById('professorForm').addEventListener('submit', function(e) {
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
        header: {
            'Content-Type': 'handler.CreateProfessorHandler'
        },
        body: JSON.stringify(professor)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        alert('Professor cadastrado com sucesso!');
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Erro ao cadastrar professor');
    });
});

}

//cadastra aluno
function cAluno(){
document.getElementById('alunoForm').addEventListener('submit', function(e) {
    e.preventDefault();

    const nome = document.getElementById('nome').value;
    const matricula = parseInt(document.getElementById('matricula').value, 10);
    const turma = document.getElementById('turma').value;

    const aluno = {
        Nome: nome,
        Matricula: matricula,
        Turma: turma,
    };

    fetch('http://localhost:8080/api/v1/aluno', {
        method: 'POST',
        header: {
            'Content-Type': 'handler.CreateAlunoHandler'
        },
        body: JSON.stringify(aluno)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        alert('Aluno cadastrado com sucesso!');
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Erro ao cadastrar aluno');
    });
});
}

//lista Professor
function listarP() {
    fetch('http://localhost:8080/api/v1/professores', {
        method: 'GET',
        header: {
            'Content-Type': 'application/json'
        },
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        const listaDiv = document.getElementById('lista');
        listaDiv.innerHTML = '';

        if (data.data && Array.isArray(data.data)) {
            data.data.forEach(professor => {
                const professorCard = document.createElement('div');
                professorCard.className = 'professor-card';
                professorCard.innerHTML = `
                    <p><strong>Identificação:</strong> ${professor.ID}</p>
                    <p><strong>Nome:</strong> ${professor.Nome}</p>
                    <p><strong>E-mail:</strong> ${professor.Email}</p>
                    <p><strong>CPF:</strong> ${professor.CPF}</p>
                    <div class="card-actions">
                        <button class="btn btn-warning" onclick="updateProfessor(${professor.ID})">Editar</button>
                        <button class="btn btn-danger" onclick="deletarProfessor(${professor.ID})">Deletar</button>
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
        document.getElementById('lista').innerHTML = '<p>Failed to load data.</p>';
    });
}

function updateProfessor(id) {
    alert(`Editar professor com ID: ${id}`);

    const nome = document.getElementById('Nome').value;
    const email = document.getElementById('Email').value;
    const cpf = document.getElementById('CPF').value;

    const professor = { Nome: nome, Email: email, CPF: cpf };

    fetch(`http://localhost:8080/api/v1/professor?id=${id}`, {
        method: 'PUT',
        header: { 'Content-Type': 'application/json' },
        body: JSON.stringify(professor)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        alert('Professor atualizado com sucesso!');
        listarP(); // Atualiza a lista de professores
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Erro ao atualizar professor');
    });
}


function deletarProfessor(id) {
    alert(`Deletar professor com ID: ${id}`);
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
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Erro ao deletar professor');
    });
}



//lista aluno

function listarA() {
    fetch('http://localhost:8080/api/v1/alunos', {
        method: 'GET',
        header: {
            'Content-Type': 'application/json'
        },
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        const listaDiv = document.getElementById('lista');
        listaDiv.innerHTML = '';

        if (data.data && Array.isArray(data.data)) {
            data.data.forEach(aluno => {
                const alunoDIV = document.createElement('div');
                alunoDIV.innerHTML = `
                    <p>ID: ${aluno.ID}</p>
                    <p>Nome: ${aluno.Nome}</p>
                    <p>Matricula: ${aluno.Matricula}</p>
                    <p>Turma: ${aluno.Turma}</p>
                `;
                listaDiv.appendChild(alunoDIV);
            });
        } else {
            listaDiv.innerHTML = '<p>No data available.</p>';
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        document.getElementById('lista').innerHTML = '<p>Failed to load data.</p>';
    });
}
