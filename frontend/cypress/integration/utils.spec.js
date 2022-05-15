export function goLogin() {
    cy.visit('http://localhost:4200');
    const showLogin = cy.get('#showLogin');
    showLogin.click();
    return showLogin;
}

export function login(email) {
    const data = email ? email : Cypress.env('email');
    goLogin();
    cy.get('#loginEmail').type(data);
    cy.get('#loginPassword').type(data);
    cy.get('#loginEnter').click();
}
