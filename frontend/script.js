function toggleSidebar() {
    const sidebar = document.getElementById('sidebar');
    sidebar.classList.toggle('open');
}

function showSection(section) {
    document.getElementById('professor-management').classList.add('hidden');
    // Adicione lógica para ocultar outras seções aqui
    if (section === 'professor') {
        document.getElementById('professor-management').classList.remove('hidden');
    }
    // Adicione lógica para exibir outras seções aqui
}

function showForm(formType) {
    const formContainer = document.getElementById('form-container');
    formContainer.innerHTML = '';

    switch (formType) {
        case 'add':
            formContainer.innerHTML = `
                <h3>Adicionar Professor</h3>
                <form id="professorForm">
                    <div class="form-group">
                        <label for="nome">Nome:</label>
                        <input type="text" class="form-control" id="nome" required>
                    </div>
                    <div class="form-group">
                        <label for="email">Email:</label>
                        <input type="email" class="form-control" id="email" required>
                    </div>
                    <div class="form-group">
                        <label for="cpf">CPF:</label>
                        <input type="text" class="form-control" id="cpf" required>
                    </div>
                    <button type="submit" class="btn btn-primary">Cadastrar</button>
                </form>
            `;
            // Adicione event listener para o formulário aqui
            document.getElementById('professorForm').addEventListener('submit', function(e) {
                e.preventDefault();
                cProfessor(); // Chama a função do arquivo cadastro.js
            });
            break;
        case 'edit':
            formContainer.innerHTML = `
                <h3>Editar Professor</h3>
                <form id="edit-professor-form">
                    <div class="form-group">
                        <label for="id">ID do Professor:</label>
                        <input type="text" id="id" name="id" class="form-control" required>
                    </div>
                    <div class="form-group">
                        <label for="name">Nome:</label>
                        <input type="text" id="name" name="name" class="form-control">
                    </div>
                    <div class="form-group">
                        <label for="subject">Disciplina:</label>
                        <input type="text" id="subject" name="subject" class="form-control">
                    </div>
                    <button type="submit" class="btn btn-warning">Editar</button>
                </form>
            `;
            break;
        case 'list':
            formContainer.innerHTML = `
                <h3>Listar Professores</h3>
                <div id="lista" class="scrollable-list">
                    <!-- Lista de professores será inserida aqui -->
                </div>
            `;
            // Adicione o listener para listar professores
            listarP(); // Certifique-se de que listarP() está disponível globalmente
            break;
        case 'delete':
            formContainer.innerHTML = `
                <h3>Apagar Professor</h3>
                <form id="delete-professor-form">
                    <div class="form-group">
                        <label for="id">ID do Professor:</label>
                        <input type="text" id="id" name="id" class="form-control" required>
                    </div>
                    <button type="submit" class="btn btn-danger">Apagar</button>
                </form>
            `;
            break;
    }
}
