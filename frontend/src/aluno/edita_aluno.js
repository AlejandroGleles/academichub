document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    if (id) {
        exibirAluno(id);
    } else {
        console.error('ID do aluno não encontrado na URL.');
    }

    const updateForm = document.getElementById('updateAlunoForm');

    if (updateForm) {
        updateForm.addEventListener('submit', (event) => {
            event.preventDefault();

            const updatedNome = document.getElementById('updateNome').value;
            const updateMatricula = document.getElementById('updateMatricula').value;
            const updateTurma = document.getElementById('updateTurma').value;

            // Validação para garantir que a matrícula seja um número inteiro
            if (!Number.isInteger(Number(updateMatricula))) {
                alert('A matrícula deve ser um número inteiro.');
                return;
            }

            const aluno = {
                Nome: updatedNome,
                Matricula: parseInt(updateMatricula, 10),
                Turma: updateTurma
            };

            atualizarAluno(aluno, id);
        });
    } else {
        console.error('Formulário de atualização não encontrado no DOM.');
    }
});


function exibirAluno(id) {
    fetch(`http://localhost:8080/api/v1/aluno?id=${id}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Erro ao buscar aluno');
        }
        return response.json();
    })
    .then(data => {
        console.log('Dados do aluno:', data);
        if (data.error) {
            alert('Erro ao exibir aluno: ' + data.error);
        } else {
         
            fillAlunoForm(data.data);
        }
    })
    .catch(error => {
        console.error('Erro na requisição:', error);
        alert('Erro ao exibir aluno');
    });
}

function fillAlunoForm(aluno) {
    const idField = document.getElementById('updateAlunoid');
    const nomeField = document.getElementById('updateNome');
    const matriculaField = document.getElementById('updateMatricula');
    const turmaField = document.getElementById('updateTurma');

    if (idField) idField.value = aluno.ID || '';
    if (nomeField) nomeField.value = aluno.Nome || '';
    if (matriculaField) matriculaField.value = aluno.Matricula || '';  // Certifique-se de que o campo de matrícula é preenchido
    if (turmaField) turmaField.value = aluno.Turma || '';
}


function atualizarAluno(aluno, id) {
    console.log('Dados a serem enviados para atualização:', aluno);

    fetch(`http://localhost:8080/api/v1/aluno?id=${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(aluno)
    })
    .then(response => {
        console.log('Resposta bruta do servidor:', response);
        if (!response.ok) {
            throw new Error('Erro ao atualizar aluno');
        }
        return response.json(); // Retorna os dados atualizados do servidor
    })
    .then(data => {
        console.log('Dados retornados pelo servidor:', data);

        if (data.error) {
            alert('Erro ao atualizar aluno: ' + data.error);
        } else {
            alert('Aluno atualizado com sucesso!');

            // Atualiza os dados exibidos no formulário após a atualização
            fillAlunoForm(data.data);

            // Atualiza o cartão do aluno na página de listagem, se existir
            atualizarCardAluno(id, data.data);
        }
    })
    .catch(error => {
        console.error('Erro na requisição:', error);
        alert('Erro ao atualizar aluno');
    });
}


function atualizarCardAluno(id, dadosAtualizados) {
    // Encontra o cartão do aluno pelo ID
    const alunoCard = document.querySelector(`.card[data-id='${id}']`);
    if (alunoCard) {
        const cardTitle = alunoCard.querySelector('.card-title');
        const cardMatricula = alunoCard.querySelector('.card-matricula'); // Alterado para uma classe específica
        const cardTurma = alunoCard.querySelector('.card-turma'); // Alterado para uma classe específica

        if (cardTitle) cardTitle.textContent = dadosAtualizados.Nome || '';
        if (cardMatricula) cardMatricula.innerHTML = `<strong>Matrícula:</strong> ${dadosAtualizados.Matricula !== undefined ? dadosAtualizados.Matricula : ''}`;
        if (cardTurma) cardTurma.innerHTML = `<strong>Turma:</strong> ${dadosAtualizados.Turma || ''}`;
    } else {
        console.error('Card do aluno não encontrado!');
    }
}

