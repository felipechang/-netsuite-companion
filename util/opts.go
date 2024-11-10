package util

func GetOptions() []string {
	return []string{
		"project: Create or select a new project to manage your NetSuite development",
		"bundle: Bundle scripts can be of type customization or configuration, allowing you to group related scripts together",
		"client: Client scripts are executed by predefined event triggers in the client browser, enabling you to customize the user interface",
		"formclient: Form Client scripts are attached to forms, allowing you to add custom logic and functionality to form submissions",
		"mapreduce: Map/Reduce scripts are designed to handle large amounts of data, making them ideal for data processing and analysis tasks",
		"massupdate: Mass update scripts allow you to programmatically perform custom updates to fields that are not available through general mass updates",
		"portlet: Portlet scripts are run on the server and are rendered in the NetSuite dashboard, providing a way to customize the dashboard with custom functionality",
		"restlet: RESTlet is a SuiteScript that you make available for other applications to call, enabling integration with external services and systems",
		"scheduled: Scheduled scripts are executed (processed) with SuiteCloud Processors, allowing you to automate tasks and processes at specific times or intervals",
		"suitelet: Suitelets are extensions of the SuiteScript API that allow you to build custom NetSuite pages and backend logic",
		"userevent: User event scripts are executed when users perform actions on records, such as create, load, update, copy, delete, or submit, enabling you to automate tasks",
		"workflowaction: Workflow action scripts are good for custom logic or managing sublist fields which are not currently available",
		"module: A module is a container for scripts, providing a way to organize and manage your code",
		"type: Holds TypeScript definitions for your scripts, providing a way to define the structure and types of your code",
	}
}
