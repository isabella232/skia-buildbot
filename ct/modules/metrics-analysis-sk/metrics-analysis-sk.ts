/**
 * @fileoverview The bulk of the Metrics Analysis page of CT.
 */

import 'elements-sk/icon/delete-icon-sk';
import 'elements-sk/icon/cancel-icon-sk';
import 'elements-sk/icon/check-circle-icon-sk';
import 'elements-sk/icon/help-icon-sk';
import 'elements-sk/toast-sk';
import '../suggest-input-sk';
import '../input-sk';
import '../patch-sk';
import '../pageset-selector-sk';
import '../task-repeater-sk';
import '../task-priority-sk';

import { $$, $ } from 'common-sk/modules/dom';
import { define } from 'elements-sk/define';
import 'elements-sk/select-sk';
import { errorMessage } from 'elements-sk/errorMessage';
import { html } from 'lit-html';

import { ElementSk } from '../../../infra-sk/modules/ElementSk';

import { InputSk } from '../input-sk/input-sk';
import { PagesetSelectorSk } from '../pageset-selector-sk/pageset-selector-sk';
import { PatchSk } from '../patch-sk/patch-sk';
import { TaskPrioritySk } from '../task-priority-sk/task-priority-sk';
import { TaskRepeaterSk } from '../task-repeater-sk/task-repeater-sk';
import {
  combineClDescriptions,
  moreThanThreeActiveTasksChecker,
} from '../ctfe_utils';
import { MetricsAnalysisAddTaskVars } from '../json';

interface ExpandableTextArea extends InputSk {
  open: boolean;
}

export class MetricsAnalysisSk extends ElementSk {
  private _triggeringTask: boolean = false;

  private _moreThanThreeActiveTasks = moreThanThreeActiveTasksChecker();

  constructor() {
    super(MetricsAnalysisSk.template);
  }

  private static template = (el: MetricsAnalysisSk) => html`

<table class=options>
  <tr>
    <td>Metric Name</td>
    <td>
      <input-sk
        id=metric_name
        label="The metric to parse the provided traces with. Eg: loadingMetric"
        class=medium-field
      ></input-sk>
    </td>
  </tr>
  <tr>
    <td>Source of Traces</td>
    <td>
      <input-sk
        id=analysis_task_id
        label="Analysis Task Id"
        class=medium-field
      ></input-sk>
      <expandable-textarea-sk
        id=custom_traces
        minRows=5
        displaytext="Specify custom list of traces"
        placeholder="Eg: trace1,trace2,trace3"
        @click=${el._toggleAnalysisTaskId}>
      </expandable-textarea-sk>
    </td>
  </tr>
  <tr>
    <td>Benchmark Arguments</td>
    <td>
      <input-sk value="--output-format=csv" id=benchmark_args class=long-field></input-sk>
      <span class=smaller-font>These will be the arguments to the analysis_metrics_ct benchmark.</span><br/>
      <span class=smaller-font><b>Note:</b> Use --run-benchmark-timeout=[secs] to specify the timeout of the run_benchmark script. 300 is the default.</span><br/>
      <span class=smaller-font><b>Note:</b> Use --max-pages-per-bot=[num] to specify the number of pages to run per bot. 50 is the default.</span>
    </td>
  </tr>
  <tr>
    <td>Field Value Column Name</td>
    <td>
      <input-sk value="avg" id=value_column_name class="medium-field"></input-sk>
      <span class="smaller-font">Which column's entries to use as field values.</span>
    </td>
  </tr>
  <tr>
    <td>
      Chromium Git patch (optional)<br/>
      Applied to Chromium ToT<br/>
    </td>
    <td>
      <patch-sk id=chromium_patch
                patchType=chromium
                @cl-description-changed=${el._patchChanged}>
      </patch-sk>
    </td>
  </tr>
  <tr>
    <td>
      Catapult Git patch (optional)<br/>
      Applied to Catapult Rev in <a href="https://chromium.googlesource.com/chromium/src/+show/HEAD/DEPS">DEPS</a>
    </td>
    <td>
      <patch-sk id=catapult_patch
                patchType=catapult
                @cl-description-changed=${el._patchChanged}>
      </patch-sk>
    </td>
  </tr>
  <tr>
    <td>Repeat this task</td>
    <td>
      <task-repeater-sk id=repeat_after_days></task-repeater-sk>
    </td>
  </tr>
  <tr>
    <td>Task Priority</td>
    <td>
      <task-priority-sk id=task_priority></task-priority-sk>
    </td>
  </tr>
  <tr>
    <td>
      Notifications CC list (optional)<br/>
      Email will be sent by ct@skia.org
    </td>
    <td>
      <input-sk value="" id=cc_list label="email1,email2,email3" class=long-field></input-sk>
    </td>
  </tr>
  <tr>
    <td>Description</td>
    <td>
      <input-sk value="" id=description label="Description is required" class=long-field></input-sk>
    </td>
  </tr>
  <tr>
    <td colspan="2" class="center">
      <div class="triggering-spinner">
        <spinner-sk .active=${el._triggeringTask} alt="Trigger task"></spinner-sk>
      </div>
      <button id=submit ?disabled=${el._triggeringTask} @click=${el._validateTask}>Queue Task
      </button>
    </td>
  </tr>
  <tr>
    <td colspan=2 class=center>
      <button id=view_history @click=${el._gotoRunsHistory}>View runs history</button>
    </td>
  </tr>
</table>
`;

  connectedCallback(): void {
    super.connectedCallback();
    this._render();
  }

  _toggleAnalysisTaskId(): void {
    const customTracesInput = $$('#custom_traces', this) as ExpandableTextArea;
    const taskInput = $$('#analysis_task_id', this) as InputSk;
    if (customTracesInput.open === taskInput.hidden) {
      // This click wasn't toggling the expandable textarea.
      return;
    }
    taskInput.hidden = customTracesInput.open;
    // We assume if someone opens the custom traces window they will be using
    // custom traces, and if they close it, they won't be.
    if (customTracesInput.open) {
      taskInput.value = '';
    } else {
      customTracesInput.value = '';
    }
    this._render();
  }

  _patchChanged(): void {
    ($$('#description', this)! as InputSk).value = combineClDescriptions(
      $('patch-sk', this).map((patch) => (patch as PatchSk).clDescription),
    );
  }

  _validateTask(): void {
    if (!$('patch-sk', this).every((patch) => (patch as PatchSk).validate())) {
      return;
    }
    if (!($$('#metric_name', this) as InputSk).value) {
      errorMessage('Please specify a metric name');
      ($$('#metric_name', this) as InputSk).focus();
      return;
    }
    if (!($$('#analysis_task_id', this) as InputSk).value && !($$('#custom_traces', this) as InputSk).value) {
      errorMessage('Please specify an analysis task id or custom traces');
      ($$('#analysis_task_id', this) as InputSk).focus();
      return;
    }
    if (!($$('#description', this) as InputSk).value) {
      errorMessage('Please specify a description');
      ($$('#description', this) as InputSk).focus();
      return;
    }
    if (this._moreThanThreeActiveTasks()) {
      return;
    }
    const confirmed = window.confirm('Proceed with queueing task?');
    if (confirmed) {
      this._queueTask();
    }
  }

  _queueTask(): void {
    this._triggeringTask = true;
    const params = {} as MetricsAnalysisAddTaskVars;
    params.metric_name = ($$('#metric_name', this) as InputSk).value;
    params.analysis_task_id = ($$('#analysis_task_id', this) as InputSk).value;
    params.custom_traces = ($$('#custom_traces', this) as PagesetSelectorSk).customPages;
    params.benchmark_args = ($$('#benchmark_args', this) as InputSk).value;
    params.value_column_name = ($$('#value_column_name', this) as InputSk).value;
    params.desc = ($$('#description', this) as InputSk).value;
    params.chromium_patch = ($$('#chromium_patch', this) as PatchSk).patch;
    params.catapult_patch = ($$('#catapult_patch', this) as PatchSk).patch;
    params.repeat_after_days = ($$('#repeat_after_days', this) as TaskRepeaterSk).frequency;
    params.task_priority = ($$('#task_priority', this) as TaskPrioritySk).priority;
    if (($$('#cc_list', this) as InputSk).value) {
      params.cc_list = ($$('#cc_list', this) as InputSk).value.split(',');
    }

    fetch('/_/add_metrics_analysis_task', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(params),
    })
      .then(() => this._gotoRunsHistory())
      .catch((e) => {
        this._triggeringTask = false;
        errorMessage(e);
      });
  }

  _gotoRunsHistory(): void {
    window.location.href = '/metrics_analysis_runs/';
  }
}

define('metrics-analysis-sk', MetricsAnalysisSk);
