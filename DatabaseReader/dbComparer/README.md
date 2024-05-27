### Assessing Damage

Okay, so now the owner decided to compare the databases. You've seen that the stolen database has modified versions of the same recipes, meaning there are several possible cases:

New cake is added or old one removed
Cooking time is different for the same cake
New ingredient is added or removed for the same cake. Important: the order of ingredients doesn't matter. Only the names are.
The count of units for the same ingredient has changed.
The unit itself for measuring the ingredient has changed.
Ingredient unit is missing or added
Quickly looking through the database, the owner also noticed that nobody bothered renaming the cakes or the ingredients (possibly, only added some new ones), so you may assume names are the same across both databases.

Your application should be runnable like this:

`~$ ./compareDB --old original_database.xml --new stolen_database.json`

It should work with both formats (JSON and XML) for original AND new database, reusing the code from Reader.

The output should look like this (the same cases explained above):

```
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake "Red Velvet Strawberry Cake"
```
