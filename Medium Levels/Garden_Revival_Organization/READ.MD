Write a function organizeGardenRevival that takes plotStatus, volunteerNames, toolInventory and returns a formatted organization report.

The function processes garden plot statuses to identify overgrown plots, assigns volunteers to them in balanced groups, and creates a tool restock notice.

Logic:

Extract plot IDs from overgrown plots using substring operations (plots with "overgrown" status)
Assign volunteers to overgrown plots using round-robin distribution (slice volunteers cyclically)
Concatenate all tool inventory items into a single restock notice
Parameters:

plotStatus ([]string): Plot status entries in format "PlotX:status"
volunteerNames ([]string): Available volunteer names
toolInventory ([]string): Tools that need restocking
Returns: Formatted report string. Format: Overgrown plots: Plot1,Plot3 Volunteer assignments: Plot1-Alice,Plot3-Bob Restock needed: shovel,rake,hose