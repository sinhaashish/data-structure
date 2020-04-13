---
title: Tree Traversal
position_number: 1.0
type:
description: Tree Travesal (PreOrder, InOrder, PostOrder)
parameters:
  - name: offset
    content: Offset the results by this amount
  - name: limit
    content: Limit the number of books returned
content_markdown: |-
  This call will return a maximum of 100 hahahaha hhhhhh jjjj books
  {: .info }

  Tree trasverasal is mainly of four kinds 
    1. PreOrder
    2. InOrder
    3. PostOrder
    4. LevelOrder

  PreOrder  : The root node is visted first then the left and then the right node.
  
  InOrder   : The left node is visted first then the root and then the right node.

  PostOrder : The left node is visted first then the right and then the root node.

  PreOrder : N L R

  InOrder  : L N R

  PostOrder: L R N


left_code_blocks:
  - code_block: |-
      $.get("http://api.myapp.com/books/", { "token": "YOUR_APP_KEY"}, function(data) {
        alert(data);
      });
    title: jQuery
    language: javascript 
  - code_block: |-
      r = requests.get("http://api.myapp.com/books/", token="YOUR_APP_KEY")
      print r.text
    title: Python
    language: python
  - code_block: |-
      var request = require("request");
      request("http://api.myapp.com/books?token=YOUR_APP_KEY", function (error, response, body) {
      if (!error && response.statusCode == 200) {
        console.log(body);
      }
    title: Node.js
    language: javascript
  - code_block: |-
      func preOrder( node *Node) {
        if node != nil {
          fmt.Println(" - mmm", node.data)
          preOrder(node.left)
          preOrder(node.right)
        }
      }

      func InOrder( node *Node) {
        if node != nil {
          fmt.Println(" - ", node.data)
          InOrder(node.left)
          InOrder(node.right)
        }
      }
    title: Go
    language: bash
right_code_blocks:
  - code_block: |2-
      [
        {
          "id": 1,
          "title": "The Hunger Games",
          "score": 4.5,
          "dateAdded": "12/12/2013"
        },
        {
          "id": 1,
          "title": "The Hunger Games",
          "score": 4.7,
          "dateAdded": "15/12/2013"
        },
      ]
    title: Response
    language: json
  - code_block: |2-
      {
        "error": true,
        "message": "Invalid offset"
      }
    title: Error
    language: json
---