package file

import (
	"bytes"
	"fmt"
	"netsuite-companion/store"
	"netsuite-companion/util"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type ClientScript struct {
	CompanyName  string
	Date         string
	Description  string
	Project      string
	UserEmail    string
	UserName     string
	ScriptName   string
	ScriptId     string
	ScriptPath   string // Path to js file
	DeploymentId string // Path to js file
}

// parseTemplate parses a template and replaces placeholders with script data
func (s *Tree) parseTemplate(script *ClientScript, name string, text string) (string, error) {
	t, err := template.New(name).Parse(text)
	if err != nil {
		return "", err
	}
	var result bytes.Buffer
	err = t.Execute(&result, script)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

// addDeploymentFiles adds deployment files for a script
func (s *Tree) addDeploymentFiles(global *store.GlobalStore, project *store.ProjectStore, scriptType string, ts string, xml string) error {
	if ts == "" && xml == "" {
		return fmt.Errorf("typescript or xml content must be set")
	}
	var fileName = util.GetInput("Enter the file name: ")
	if fileName == "" {
		return fmt.Errorf("file name cannot be empty")
	}
	fileNameParsed := strings.ReplaceAll(strings.ToLower(fileName), " ", "_")
	var description = util.GetInput("Enter the file description: ")
	if description == "" {
		return fmt.Errorf("file description cannot be empty")
	}
	filePattern := fmt.Sprintf("%s_%s", global.VendorPrefix, fileNameParsed)
	projectPath := filepath.Join("SuiteScripts", global.VendorName, project.Current)
	scriptPath := filepath.Join(projectPath, fmt.Sprintf("%s_%s.js", filePattern, scriptType))

	clientScript := &ClientScript{
		CompanyName:  global.VendorName,
		Date:         time.Now().Format("01/02/2006"),
		Description:  description,
		Project:      project.Current,
		UserEmail:    global.AuthorEmail,
		UserName:     global.AuthorName,
		ScriptName:   fileName,
		ScriptId:     fmt.Sprintf(`customscript_%s`, filePattern),
		ScriptPath:   fmt.Sprintf(`\%s`, scriptPath),
		DeploymentId: fmt.Sprintf(`customdeploy_%s`, filePattern),
	}
	if ts != "" {
		parsedTS, err := s.parseTemplate(clientScript, scriptType, ts)
		if err != nil {
			return err
		}
		destinationTS := filepath.Join(
			filepath.Join(s.dirname, "src", "FileCabinet", projectPath),
			fmt.Sprintf("%s_%s.ts", filePattern, scriptType),
		)
		err = s.createFile(destinationTS, parsedTS)
		if err != nil {
			return err
		}
	}
	if xml != "" {
		parsedXML, err := s.parseTemplate(clientScript, scriptType, xml)
		if err != nil {
			return err
		}
		destinationXML := filepath.Join(
			filepath.Join(s.dirname, "src", "Objects"),
			fmt.Sprintf("%s_%s.xml", filePattern, scriptType),
		)
		err = s.createFile(destinationXML, parsedXML)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateBundle creates a bundle script
func (s *Tree) CreateBundle(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "bundle", `import {EntryPoints} from "N/types";
import onAfterInstallContext = EntryPoints.BundleInstallation.onAfterInstallContext;
import onAfterUpdateContext = EntryPoints.BundleInstallation.onAfterUpdateContext;
import onBeforeInstallContext = EntryPoints.BundleInstallation.onBeforeInstallContext;
import onBeforeUpdateContext = EntryPoints.BundleInstallation.onBeforeUpdateContext;
import onBeforeUninstallContext = EntryPoints.BundleInstallation.onBeforeUninstallContext;

/**
 * Bundle Installation script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType BundleInstallationScript
 */

/** afterInstall event handler */
export let afterInstall: EntryPoints.BundleInstallation.afterInstall = (context: onAfterInstallContext) => {
    // Enter code here
};

/** afterUpdate event handler */
export let afterUpdate: EntryPoints.BundleInstallation.afterUpdate = (context: onAfterUpdateContext) => {
    // Enter code here
};

/** beforeInstall event handler */
export let beforeInstall: EntryPoints.BundleInstallation.beforeInstall = (context: onBeforeInstallContext) => {
    // Enter code here
};

/** beforeUninstall event handler */
export let beforeUninstall: EntryPoints.BundleInstallation.beforeUninstall = (context: onBeforeUninstallContext) => {
    // Enter code here
};

/** beforeUpdate event handler */
export let beforeUpdate: EntryPoints.BundleInstallation.beforeUpdate = (context: onBeforeUpdateContext) => {
    // Enter code here
};
`, ``)
	if err != nil {
		return err
	}
	return nil
}

// CreateClient creates a client script
func (s *Tree) CreateClient(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "client", `import {EntryPoints} from "N/types";

/**
 * Client script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType ClientScript
 */

/** pageInit event handler */
export let pageInit: EntryPoints.Client.pageInit = (context: EntryPoints.Client.pageInitContext) => {
    // Enter code here
};

/** validateField event handler */
export let validateField: EntryPoints.Client.validateField = (context: EntryPoints.Client.validateFieldContext) => {
    // Enter code here
};

/** fieldChanged event handler */
export let fieldChanged: EntryPoints.Client.fieldChanged = (context: EntryPoints.Client.fieldChangedContext) => {
    // Enter code here
};

/** postSourcing event handler */
export let postSourcing: EntryPoints.Client.postSourcing = (context: EntryPoints.Client.postSourcingContext) => {
    // Enter code here
};

/** lineInit event handler */
export let lineInit: EntryPoints.Client.lineInit = (context: EntryPoints.Client.lineInitContext) => {
    // Enter code here
};

/** validateLine event handler */
export let validateLine: EntryPoints.Client.validateLine = (context: EntryPoints.Client.validateLineContext) => {
    // Enter code here
};

/** validateInsert event handler */
export let validateInsert: EntryPoints.Client.validateInsert = (context: EntryPoints.Client.validateInsertContext) => {
    // Enter code here
};

/** validateDelete event handler */
export let validateDelete: EntryPoints.Client.validateDelete = (context: EntryPoints.Client.validateDeleteContext) => {
    // Enter code here
};

/** sublistChanged event handler */
export let sublistChanged: EntryPoints.Client.sublistChanged = (context: EntryPoints.Client.sublistChangedContext) => {
    // Enter code here
};

/** saveRecord event handler */
export let saveRecord: EntryPoints.Client.saveRecord = (context: EntryPoints.Client.saveRecordContext) => {
    // Enter code here
};
`, `<clientscript scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <notifyuser>F</notifyuser>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
</clientscript>`)
	if err != nil {
		return err
	}
	return nil
}

// CreateFormClient creates a form client script
func (s *Tree) CreateFormClient(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "formclient", `import {EntryPoints} from "N/types";

/**
 * Form client script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType ClientScript
 */

/** pageInit event handler */
export let pageInit: EntryPoints.Client.pageInit = (context: EntryPoints.Client.pageInitContext) => {
    // Enter code here
};

/** validateField event handler */
export let validateField: EntryPoints.Client.validateField = (context: EntryPoints.Client.validateFieldContext) => {
    // Enter code here
};

/** fieldChanged event handler */
export let fieldChanged: EntryPoints.Client.fieldChanged = (context: EntryPoints.Client.fieldChangedContext) => {
    // Enter code here
};

/** postSourcing event handler */
export let postSourcing: EntryPoints.Client.postSourcing = (context: EntryPoints.Client.postSourcingContext) => {
    // Enter code here
};

/** lineInit event handler */
export let lineInit: EntryPoints.Client.lineInit = (context: EntryPoints.Client.lineInitContext) => {
    // Enter code here
};

/** validateLine event handler */
export let validateLine: EntryPoints.Client.validateLine = (context: EntryPoints.Client.validateLineContext) => {
    // Enter code here
};

/** validateInsert event handler */
export let validateInsert: EntryPoints.Client.validateInsert = (context: EntryPoints.Client.validateInsertContext) => {
    // Enter code here
};

/** validateDelete event handler */
export let validateDelete: EntryPoints.Client.validateDelete = (context: EntryPoints.Client.validateDeleteContext) => {
    // Enter code here
};

/** sublistChanged event handler */
export let sublistChanged: EntryPoints.Client.sublistChanged = (context: EntryPoints.Client.sublistChangedContext) => {
    // Enter code here
};

/** saveRecord event handler */
export let saveRecord: EntryPoints.Client.saveRecord = (context: EntryPoints.Client.saveRecordContext) => {
    // Enter code here
};
`, ``)
	if err != nil {
		return err
	}
	return nil
}

// CreateMapReduce creates a map/reduce script
func (s *Tree) CreateMapReduce(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "mapreduce", `import {EntryPoints} from "N/types";

/**
 * Map/Reduce script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType MapReduceScript
 */

/** getInputData event handler */
export let getInputData: EntryPoints.MapReduce.getInputData = (context: EntryPoints.MapReduce.getInputDataContext) => {
    // Enter code here
};

/** map event handler */
export let map: EntryPoints.MapReduce.map = (context: EntryPoints.MapReduce.mapContext) => {
    // Enter code here
};

/** reduce event handler */
export let reduce: EntryPoints.MapReduce.reduce = (context: EntryPoints.MapReduce.reduceContext) => {
    // Enter code here
};

/** summarize event handler */
export let summarize: EntryPoints.MapReduce.summarize = (summary: EntryPoints.MapReduce.summarizeContext) => {
    // Enter code here
};
`, `<mapreducescript scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
</mapreducescript>
`)
	if err != nil {
		return err
	}
	return nil
}

// CreateMassUpdate creates a mass update script
func (s *Tree) CreateMassUpdate(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "massupdate", `import {EntryPoints} from "N/types";

/**
 * Mass Update script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType MassUpdateScript
 */

/** each event handler */
export let each: EntryPoints.MassUpdate.each = (params: EntryPoints.MassUpdate.eachContext) => {
    // Enter code here
};
`, `<massupdatescript scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <notifyuser>F</notifyuser>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
</massupdatescript>`)
	if err != nil {
		return err
	}
	return nil
}

// CreatePortlet creates a portlet script
func (s *Tree) CreatePortlet(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "portlet", `import {EntryPoints} from "N/types";

/**
 * Portlet script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType Portlet
 */

/** render event handler */
export let render: EntryPoints.Portlet.render = (params: EntryPoints.Portlet.renderContext) => {
    // Enter code here
};
`, `<portlet scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <notifyuser>F</notifyuser>
  <portlettype>HTML</portlettype>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
  <scriptdeployments>
    <scriptdeployment scriptid="{{.DeploymentId}}">
      <allemployees>T</allemployees>
      <allpartners>F</allpartners>
      <allroles>F</allroles>
      <audslctrole></audslctrole>
      <dashboardapp>F</dashboardapp>
      <isdeployed>T</isdeployed>
      <loglevel>ERROR</loglevel>
      <runasrole></runasrole>
      <status>RELEASED</status>
      <title>{{.ScriptName}}</title>
    </scriptdeployment>
  </scriptdeployments>
</portlet>
`)
	if err != nil {
		return err
	}
	return nil
}

// CreateRestlet creates a restlet script
func (s *Tree) CreateRestlet(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "restlet", `import {EntryPoints} from "N/types";

/** RESTlet standard return */
type RestReturn = string | object;

/**
 * RESTlet script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType Restlet
 */

/** GET event handler */
const get: EntryPoints.RESTlet.get = (requestParams: object): RestReturn => {
    // Enter code here
};

/** POST event handler */
const post: EntryPoints.RESTlet.post = (requestBody: object): RestReturn => {
    // Enter code here
};

/** PUT event handler */
const put: EntryPoints.RESTlet.put = (requestBody: object): RestReturn => {
    // Enter code here
};

/** DELETE event handler */
const remove: EntryPoints.RESTlet.delete_ = (requestParams: object): RestReturn => {
    // Enter code here
};

export = {
    ["get"]: get,
    ["post"]: post,
    ["put"]: put,
    ["delete"]: remove,
};
`, `<restlet scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <notifyuser>F</notifyuser>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
  <scriptdeployments>
    <scriptdeployment scriptid="{{.DeploymentId}}">
      <allemployees>T</allemployees>
      <allpartners>F</allpartners>
      <allroles>F</allroles>
      <audslctrole></audslctrole>
      <isdeployed>T</isdeployed>
      <loglevel>ERROR</loglevel>
      <status>RELEASED</status>
      <title>{{.ScriptName}}</title>
    </scriptdeployment>
  </scriptdeployments>
</restlet>`)
	if err != nil {
		return err
	}
	return nil
}

// CreateScheduled creates a scheduled script
func (s *Tree) CreateScheduled(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "scheduled", `import {EntryPoints} from "N/types";

/**
 * Scheduled script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType ScheduledScript
 */

/** execute event handler */
export let execute: EntryPoints.Scheduled.execute = (context: EntryPoints.Scheduled.executeContext) => {
    // Enter code here
};
`, `<scheduledscript scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
  <scriptdeployments>
    <scriptdeployment scriptid="{{.DeploymentId}}">
      <isdeployed>T</isdeployed>
      <loglevel>DEBUG</loglevel>
      <status>NOTSCHEDULED</status>
      <title>{{.ScriptName}}</title>
      <recurrence>
        <single>
          <repeat></repeat>
          <startdate>2020-01-01</startdate>
          <starttime>23:00:00Z</starttime>
        </single>
      </recurrence>
    </scriptdeployment>
  </scriptdeployments>
</scheduledscript>`)
	if err != nil {
		return err
	}
	return nil
}

// CreateSuitelet creates a suitelet script
func (s *Tree) CreateSuitelet(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "suitelet", `import {EntryPoints} from "N/types";

/**
 * Suitelet script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType Suitelet
 */

/** onRequest event handler */
export let onRequest: EntryPoints.Suitelet.onRequest = (context: EntryPoints.Suitelet.onRequestContext) => {
    // Enter code here
};
`, `<suitelet scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <notifyuser>F</notifyuser>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
  <scriptdeployments>
    <scriptdeployment scriptid="{{.DeploymentId}}">
      <allemployees>T</allemployees>
      <allpartners>F</allpartners>
      <allroles>F</allroles>
      <audslctrole></audslctrole>
      <eventtype></eventtype>
      <isdeployed>T</isdeployed>
      <isonline>F</isonline>
      <loglevel>ERROR</loglevel>
      <runasrole>ADMINISTRATOR</runasrole>
      <status>RELEASED</status>
      <title>{{.ScriptName}}</title>
    </scriptdeployment>
  </scriptdeployments>
</suitelet>`)
	if err != nil {
		return err
	}
	return nil
}

// CreateUserEvent creates a user event script
func (s *Tree) CreateUserEvent(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "userevent", `import {EntryPoints} from "N/types";

/**
 * User Event script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType UserEventScript
 */

/** beforeLoad event handler */
export let beforeLoad: EntryPoints.UserEvent.beforeLoad = (context: EntryPoints.UserEvent.beforeLoadContext) => {
    // Enter code here
};

/** beforeSubmit event handler */
export let beforeSubmit: EntryPoints.UserEvent.beforeSubmit = (context: EntryPoints.UserEvent.beforeSubmitContext) => {
    // Enter code here
};

/** afterSubmit event handler */
export let afterSubmit: EntryPoints.UserEvent.afterSubmit = (context: EntryPoints.UserEvent.afterSubmitContext) => {
    // Enter code here
};
`, `<usereventscript scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <notifyuser>F</notifyuser>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
</usereventscript>`)
	if err != nil {
		return err
	}
	return nil
}

// CreateWorkflowAction creates a workflow action script
func (s *Tree) CreateWorkflowAction(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "workflowaction", `import {EntryPoints} from "N/types";

/**
 * Workflow script file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NScriptName {{.ScriptName}}
 * @NScriptId {{.ScriptId}}
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 * @NScriptType WorkflowActionScript
 */

/** onAction event handler */
export let onAction: EntryPoints.WorkflowAction.onAction = (context: EntryPoints.WorkflowAction.onActionContext) => {
    // Enter code here
};
`, `<workflowactionscript scriptid="{{.ScriptId}}">
  <description>{{.Description}}</description>
  <isinactive>F</isinactive>
  <name>{{.ScriptName}}</name>
  <notifyadmins>F</notifyadmins>
  <notifyemails></notifyemails>
  <notifyowner>T</notifyowner>
  <notifyuser>F</notifyuser>
  <returnrecordtype>-4</returnrecordtype>
  <returntype>SELECT</returntype>
  <scriptfile>[{{.ScriptPath}}]</scriptfile>
</workflowactionscript>`)
	if err != nil {
		return err
	}
	return nil
}

// CreateModule creates a module file
func (s *Tree) CreateModule(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "type", `/**
 * Type declaration file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 */

`, ``)
	if err != nil {
		return err
	}
	return nil
}

// CreateType creates a type file
func (s *Tree) CreateType(global *store.GlobalStore, project *store.ProjectStore) error {
	err := s.addDeploymentFiles(global, project, "module", `/**
 * Module file
 *
 * WARNING:
 * TypeScript generated file, do not edit directly
 * source files are located in the repository
 *
 * @project: {{.Project}}
 * @description: {{.Description}}
 *
 * @copyright {{.Date}} {{.CompanyName}}
 * @author {{.UserName}} {{.UserEmail}}
 *
 * @NApiVersion 2.x
 * @NModuleScope SameAccount
 */

export {};
`, ``)
	if err != nil {
		return err
	}
	return nil
}
