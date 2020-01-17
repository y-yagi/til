SELECT Name AS Customers FROM Customers c WHERE c.Id NOT IN (SELECT CustomerID FROM Orders)
