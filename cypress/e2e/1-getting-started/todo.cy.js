
describe('Login Component', () => {


  it('should test all of main fucntionality', () => {

    cy.visit('http://localhost:4200/main')
    
    cy.get('button').contains('Link Spotify').click(); // click the Redirect to Spotify API button


  
    const checkUrl = (counter) => {
      if (counter > 60) {
        throw new Error('Timed out waiting for redirect to localhost');
      }
  
      cy.url().then(url => {
        if (!url.includes('http://localhost:4200/')) {
  
          cy.wait(1000); // wait 1 second before checking the URL again
          checkUrl(counter + 1); // call the function recursively with an incremented counter
        } else {
          cy.url().should('include', 'http://localhost:4200/'); // make sure the URL has changed to the localhost URL
        }
      });
    };
  
    checkUrl(0); // start the recursive function with a counter of 0

    cy.get('button').contains('Login').click();

    cy.wait(10000); // wait for 2 seconds to allow the select element to load
    

    cy.get('[data-cy=nameSelect]').select('akshatpant3002');
    cy.get('[data-cy=addFriendButton]').click();
    
    

  });
  
  // it('should handle auth', () => {
  //   cy.get('button').contains('Login').click(); // click the handle auth button
  
  // });

  
    // it('should add a friend to the list', () => {
    //   const friendName = 'John';
    //   cy.get('#nameSelect').type(friendName); // type the friend name in the input field
    //   cy.get('button').contains('Add Friend').click(); // click the "Add Friend" button
    //   cy.get('ul li').last().should('have.text', friendName); // verify that the last item in the list is the new friend
    // });

  // it('should add multiple friends and select a Random Friend', () => {
  //   const friendNames = ['John', 'Sarah', 'Mike', 'Akshat', 'Aryaan Verma', 'Vasu Sorathia', 'Drake', 'Kevin']; // set an array of friend names to use in the test
  //   cy.wrap(friendNames).each((friendName) => { // loop through the friend names array using Cypress' .each() method
  //     cy.get('input').type(friendName); // type the friend name in the input field
  //     cy.get('button').contains('Add Friend').click(); // click the Add Friend button
  //     cy.get('ul').contains(friendName); // make sure the friend name is displayed in the list of friends
  //     cy.get('input').clear(); // clear the input field for the next friend name
  //   });
  //   const randomIndex = Math.floor(Math.random() * friendNames.length);
  //   const selectedFriend = friendNames[randomIndex];
    
  //   // Select the friend from the dropdown
  //   cy.get('.friend-selector').within(() => {
  //     cy.get('select').select(selectedFriend);
  //   });
    
  //   // Verify that the selected friend is displayed in the list
  //   cy.get('ul').contains(selectedFriend).should('be.visible');
  // });


  //   it('should navigate to home page', () => {
  //     cy.get('a').contains('Home').click(); // click the "Home" link
  //     cy.url().should('include', '/home'); // make sure the URL has changed to the home page URL
  //   });
  
  //   it('should navigate to login page', () => {
  //     cy.get('a').contains('Login').click(); // click the "Login" link
  //     cy.url().should('include', '/login'); // make sure the URL has changed to the login page URL
  //   });

  // it('should display songs on submit', () => {
  //   cy.get('button').contains('Display Songs').click(); // click the Display Songs button
  //   cy.get('.song-card').should('have.length', 1); // make sure at least one song card is displayedd
  // });
});
