Feature: Save issue to database
  In order to analyze issue data
  As a repository administrator
  I need to keep updated issue data in an SQL database

  Scenario: An issue with ID 1 is opened
    Given there are no issues in the database with ID 1
    When an issue with ID 1 is opened
    Then there should be an issue with ID 1 in the database

  Scenario: An issue with ID 2 is edited
    Given there is an issue with ID 2 in the database
    And its body is "Original issue text"
    When the body of issue with ID 2 is changed to "This is the new text"
    And this change is made at 2021-08-01T13:00:00Z
    Then there should be an issue with ID 2 in the database
    And its body should be "This is the new text"
    And its update timestamp should be "2021-08-01T13:00:00Z"

  Scenario: An issue with ID 3, with no labels or comments, is deleted
    Given there is an issue with ID 3 in the database
    And it has no labels
    And it has no comments
    When the issue with ID 3 is deleted
    Then there should not be an issue with ID 3 in the database

  Scenario: An issue with ID 4, with labels but no comments, is deleted
    Given there is an issue with ID 4 in the database
    And it has a label with ID 1 and text "bug"
    And it has a label with ID 2 and text "critical"
    And it has no comments
    When the issue with ID 4 is deleted
    Then there should not be an issue with ID 4 in the database
    And the label with ID 1 should not be applied to the issue with ID 4
    And the label with ID 2 should not be applied to the issue with ID 4

  Scenario: An issue with ID 5, with no labels but with comments, is deleted
    Given there is an issue with ID 5 in the database
    And it has no labels
    And it has a comment with ID 1
    When the issue with ID 5 is deleted
    Then there should not be an issue with ID 5 in the database
    And there should not be a comment with ID 1 in the database

  Scenario: An issue with ID 6, with label and comments, is deleted
    Given there is an issue with ID 6 in the database
    And it has a label with ID 3 and the text "feature"
    And it has a comment with ID 2
    When the issue with ID 6 is deleted
    Then there should not be an issue with ID 6 in the database
    And the label with ID 3 should not be applied to the issue with ID 6
    And there should not be a comment with ID 2 in the database

  Scenario: An issue with ID 7 is closed
    Given there is an issue with ID 7 in the database
    And its state is "open"
    And it has not been closed yet (its "closed_at" field is null)
    When the issue with ID 7 is closed at 2021-08-01T14:00:00Z
    Then there should be an issue with ID 7 in the database
    And its state should be "closed"
    And its close timestamp should be "2021-08-01T14:00:00Z"

  Scenario: An issue with ID 8 is reopened
    Given there is an issue with ID 8 in the database
    And its state is "closed"
    When the issue with ID 8 is reopened
    Then there should be an issue with ID 8 in the database
    And its state should be "reopened"

  Scenario: An issue with ID 9 is labeled
    Given there is an issue with ID 9 in the database
    And it has a label with ID 4 and the text "analysis"
    When the issue with ID 9 is labeled
    And the label applied has ID 5 and text "important"
    Then there should be an issue with ID 9 in the database
    And there should be a label with ID 5 and text "important" in the database
    And the label with ID 6 should be applied to the issue with ID 9

  Scenario: An issue with ID 10 is unlabeled
    Given there is an issue with ID 10 in the database
    And it has a label with ID 6 and the text "documentation"
    When the label with ID 6 is removed from issue with ID 10
    Then there should be an issue with ID 10 in the database
    And the label with ID 6 should not be applied to the issue with ID 10

  Scenario: A comment is created
    Given there is an issue with ID 11 in the database
    And it has no comments
    And there are no comments with ID 1 in the database
    When a comment with ID 1 is created in the issue with ID 11
    Then there should be a comment with ID 1 in the database
    And the comment with ID 1 should have "issue_id" 11

  Scenario: A comment is edited
    Given there is an issue with ID 12 in the database
    And it has a comment with ID 2
    And its body is "This is the original comment"
    When the body of the comment with ID 2 is changed to "This is the new comment"
    And this change is made at 2021-08-01T15:00:00Z
    Then there should be a comment with ID 2
    And it should have "issue_id" 12
    And its body should be "This is the new comment"
    And its update timestamp should be "2021-08-01T15:00:00Z"
  
  Scenario: A comment is deleted
    Given there is an issue with ID 13 in the database
    And it has a comment with ID 3
    When the comment with ID 3 is deleted
    Then there should not be a comment with ID 3 in the database
