import { login } from '../utils.spec';

describe('Create group', () => {
    let email;

    beforeEach(() => {
        email = Cypress.env('email');
    });

    it('should create a private group with a user without permissions', () => {
        login();
        cy.get('#groupsMenuElement').click();
        cy.get('#createGroupsButton').click();
        const randomNumber = new Date().getTime();
        const group = `group${randomNumber}`;
        Cypress.env('group', group);
        const groupName = cy.get('#groupName');
        groupName.type(group).should('have.value', group);
        cy.get('#groupPublic').click().should('not.have.class', 'ant-checkbox-wrapper-checked');
        cy.get('#groupSend').click();
        groupName.should('not.exist');
    });
});
